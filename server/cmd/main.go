package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	fmt.Println("hello world")
	e := echo.New()

	// db := datastore.NewDB()
	// taskRepository := datastore.NewTaskRepository(db)
	// taskApplicationService := application.NewTaskApplicationService(taskRepository)
	// taskHandler := handlers.NewTaskHandler(taskApplicationService)

	fmt.Println("CORS set")
	e.Use(middleware.CORS())

	taskHandler := InitializeTaskHandler()
	taskHandler.Routes(e)
	e.Start(":1323")
}
