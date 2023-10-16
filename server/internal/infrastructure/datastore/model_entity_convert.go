package datastore

import (
	"github.com/google/uuid"
	"github.com/morninng/z-go-react-monorepo/internal/domain/entity"
	models "github.com/morninng/z-go-react-monorepo/internal/infrastructure/model"
)

func convertTask(m *models.Task) entity.Task {
	return entity.Task{
		ID:      uuid.MustParse(m.ID),
		Name:    m.Name,
		Content: m.Content,
	}
}
