package entity

import (
	"github.com/google/uuid"
)

type Tasks []Task

type Task struct {
	ID      uuid.UUID
	Name    string
	Content string
}
