package repository

import (
	"context"

	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
)

type TaskRepository interface {
	GetAll(ctx context.Context) (*entity.Tasks, error)
}
