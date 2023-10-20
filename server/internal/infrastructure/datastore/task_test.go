package datastore

import (
	"context"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
)

func Test_ShouldGetAllTasks(t *testing.T) {

	ctx := context.Background()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expectedTasks := entity.Tasks{
		entity.Task{
			ID:      uuid.MustParse("6a065134-318e-4689-bf6f-1c5f0234b32e"),
			Name:    "sss",
			Content: "sdddd",
		},
		// 他のタスクも同様に追加
	}

	mock.ExpectQuery("SELECT \"tasks\"\\..* FROM \"tasks\"").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "content"}).
			AddRow("6a065134-318e-4689-bf6f-1c5f0234b32e", "sss", "sdddd"))
	taskRepository := NewTaskRepository(db)

	// モック化されたDBを用いてテスト対象関数を実行
	resAllTasks, err := taskRepository.GetAll(ctx)
	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 使用されたモックDBが期待通りの値を持っているかを検証
	if !reflect.DeepEqual(resAllTasks, &expectedTasks) {
		t.Errorf("unexpected result. Expected: %+v, Actual: %+v", expectedTasks, resAllTasks)
	}
}
