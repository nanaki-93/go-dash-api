package validator

import (
	"bytes"
	"encoding/json"
	"github.com/kaptinlin/jsonschema"
	"github.com/labstack/echo/v4"
	"go-dash-api/internal/service"
	"io"
	"log"
	"net/http"
)

type RequestValidator struct {
	scSrv *service.SchemaService
}

func NewJsonValidator(schemaSrv *service.SchemaService) *RequestValidator {
	return &RequestValidator{scSrv: schemaSrv}
}

func (cv *RequestValidator) Validate(i interface{}) error {
	c := i.(echo.Context)
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request().Body)
	}

	// Restore the io.ReadCloser to its original state
	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Continue to use the Body, like Binding it to a struct:
	bodyEntity := make(map[string]interface{})
	err := c.Bind(&bodyEntity)
	if err != nil {
		return err
	}

	c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	tableSchema, err := cv.scSrv.GetSchema(bodyEntity["collection"].(string))
	log.Println(tableSchema)
	if err != nil {
		log.Fatal(err)
		return err
	}

	compiler := jsonschema.NewCompiler()
	schema, err := compiler.Compile([]byte(tableSchema.Structure))
	if err != nil {
		log.Fatalf("Failed to compile schema: %v", err)
	}

	result := schema.Validate(bodyEntity)
	if !result.IsValid() {
		details, _ := json.MarshalIndent(result.ToList(), "", "  ")
		return echo.NewHTTPError(http.StatusBadRequest, string(details))
	}
	return nil
}
