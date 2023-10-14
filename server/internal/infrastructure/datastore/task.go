package datastore

import (
	"context"
	"database/sql"

	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	"github.com/morninng/z-go-react-monorepo/internal/domain/repository"
	models "github.com/morninng/z-go-react-monorepo/internal/infrastructure/model"
)

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{db: db}
}

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) GetAll(c context.Context) (*entity.Tasks, error) {

	tasks, err := models.Tasks().All(c, t.db)
	if err != nil {
		return nil, err
	}
	result := make(entity.Tasks, len(tasks))
	for i, task := range tasks {
		ta := convertTask(task)
		result[i] = ta
	}
	return &result, nil
}
