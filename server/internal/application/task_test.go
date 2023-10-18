package application

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	"github.com/morninng/z-go-react-monorepo/internal/domain/repository_mock"

	"github.com/golang/mock/gomock"
)

func TestGetAllTasks(t *testing.T) {

	ctx := context.TODO()
	allTasks := entity.Tasks{
		{ID: uuid.New(), Name: "aaa", Content: "jj"},
		{ID: uuid.New(), Name: "bbba", Content: "jj"},
	}

	// repository mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTaskRepsitory := repository_mock.NewMockTaskRepository(mockCtrl)

	mockTaskRepsitory.
		EXPECT().
		GetAll(ctx).
		Return(allTasks, nil)

	taskApplicationService := NewTaskApplicationService(mockTaskRepsitory)

	resAllTasks, _ := taskApplicationService.GetAll()
	if !reflect.DeepEqual(resAllTasks, allTasks) {
		t.Error("tasks are not equal")
	}

}
