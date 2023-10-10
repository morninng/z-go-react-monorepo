package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morninng/z-go-react-monorepo/internal/dto"
)

type TaskHandler struct {
}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (h *TaskHandler) Routes(e *echo.Echo) {
	fmt.Println("member route")

	e.GET("/tasks", h.getTasks)

}

func (h *TaskHandler) getTasks(c echo.Context) error {

	response := make(dto.Tasks, 3)

	return c.JSON(http.StatusOK, &response)
}
