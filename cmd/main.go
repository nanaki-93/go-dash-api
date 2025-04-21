package main

import (
	"go-dash-api/internal/config"
)

func main() {

	e := config.InitServer()

	e.Logger.Fatal(e.Start(":1323"))
}
