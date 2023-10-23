package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/morninng/z-go-react-monorepo/internal/application"
	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestGetTask(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetPath("/tasks/:email")
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	allTasks := entity.Tasks{
		{ID: uuid1, Name: "aaa", Content: "jj"},
		{ID: uuid2, Name: "bbba", Content: "jjs"},
	}
	tasksJson := fmt.Sprintf("[{\"constent\":\"jj\",\"id\":\"%s\",\"name\":\"aaa\"},{\"constent\":\"jjs\",\"id\":\"%s\",\"name\":\"bbba\"}]\n", uuid1.String(), uuid2.String())

	// repository mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTaskApplication := application.NewMockTaskApplicationService(mockCtrl)

	mockTaskApplication.
		EXPECT().
		GetAll(gomock.Any()).
		Return(&allTasks, nil)

	taskHandler := NewTaskHandler(mockTaskApplication)

	// Assertions
	if assert.NoError(t, taskHandler.getTasks(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, tasksJson, rec.Body.String())
	}
}
