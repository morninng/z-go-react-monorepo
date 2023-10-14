//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package application

import (
	"context"

	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	"github.com/morninng/z-go-react-monorepo/internal/domain/repository"
	"golang.org/x/sync/errgroup"
)

type (
	TaskApplicationService interface {
		GetAll(c context.Context) (*entity.Tasks, error)
	}

	taskApplicationService struct {
		taskRepo repository.TaskRepository
	}
)

func NewTaskApplicationService(taskRepo repository.TaskRepository) TaskApplicationService {
	return &taskApplicationService{taskRepo: taskRepo}
}

func (t *taskApplicationService) GetAll(c context.Context) (*entity.Tasks, error) {
	// op := "taskApplicationService.Get"
	g, ctx := errgroup.WithContext(c)
	tasksChan := make(chan *entity.Tasks, 1)

	g.Go(func() error {
		defer close(tasksChan)
		tasks, err := t.taskRepo.GetAll(ctx)
		if err != nil {
			return err
		}
		tasksChan <- tasks
		return nil
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return <-tasksChan, nil
}
