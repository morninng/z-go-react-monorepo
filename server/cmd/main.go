package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {

	fmt.Println("hello world")
	e := echo.New()

	taskHandler := InitializeTaskHandler()
	taskHandler.Routes(e)
	e.Start(":1323")
}
