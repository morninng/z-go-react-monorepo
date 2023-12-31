package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/morninng/z-go-react-monorepo/internal/application"
	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	"github.com/morninng/z-go-react-monorepo/internal/dto"
)

type TaskHandler struct {
	taskAppSrv application.TaskApplicationService
}

func NewTaskHandler(taskAppSrv application.TaskApplicationService) *TaskHandler {
	return &TaskHandler{taskAppSrv: taskAppSrv}
}

func (h *TaskHandler) Routes(e *echo.Echo) {
	fmt.Println("member route")

	e.GET("/tasks", h.getTasks)
	e.POST("/tasks", h.postTasks)

}

func (h *TaskHandler) getTasks(c echo.Context) error {
	ctx := c.Request().Context()

	tasks, err := h.taskAppSrv.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	response := make(dto.Tasks, len(*tasks))
	for i, task := range *tasks {
		response[i] = dto.Task{
			Id:       task.ID,
			Name:     task.Name,
			Constent: task.Content,
		}
	}

	return c.JSON(http.StatusOK, &response)
}

func (h *TaskHandler) postTasks(c echo.Context) error {
	ctx := c.Request().Context()

	dtoTasks := &dto.Tasks{}
	if err := c.Bind(dtoTasks); err != nil {
		fmt.Println("ssss", err)
		return c.JSON(http.StatusBadRequest, dto.BadRequestError{
			Code:    dto.BadRequestErrorCodeEnumInvalidParameter,
			Message: "sss",
		})
	}

	cmdParms := make(entity.Tasks, len(*dtoTasks))
	for i, task := range *dtoTasks {
		cmdParms[i] = entity.Task{
			ID:      uuid.UUID(task.Id),
			Name:    task.Name,
			Content: task.Constent,
		}
	}

	postCmd := application.TaskSaveCommand{
		Tasks: cmdParms,
	}

	err := h.taskAppSrv.PostMany(ctx, postCmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
