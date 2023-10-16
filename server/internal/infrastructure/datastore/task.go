package datastore

import (
	"context"
	"database/sql"
	"fmt"

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

func (t *taskRepository) CreateMany(c context.Context, tasks_entities entity.Tasks) error {

	// tasks_models := make([]models.Task, len(tasks_entities))

	tx, err := t.db.Begin()
	if err != nil {
		// エラーハンドリング
		fmt.Println("ssss", err)
		return err
	}
	defer tx.Rollback()

	for _, task := range tasks_entities {
		_, err := tx.ExecContext(
			c,
			"INSERT INTO tasks (id, name, content) VALUES ($1, $2, $3)",
			task.ID,
			task.Name,
			task.Content,
		)
		if err != nil {
			// エラーハンドリング
			fmt.Println("ttt", err)
			return err
		}

	}

	if err := tx.Commit(); err != nil {
		// エラーハンドリング
		fmt.Println("uuu", err)
		return err
	}

	return nil
}
