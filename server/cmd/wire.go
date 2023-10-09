//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morninng/z-go-react-monorepo/internal/interfaces/handlers"
)

func InitializeTaskHandler() (h *handlers.TaskHandler) {
	wire.Build(
		handlers.SuperSet,
	)
	return
}
