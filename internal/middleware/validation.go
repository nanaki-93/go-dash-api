package middle

import (
	"github.com/labstack/echo/v4"
	"log"
)

func Validation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Validation")
		err := c.Validate(c)
		if err != nil {
			return err
		}

		return next(c)
	}
}
