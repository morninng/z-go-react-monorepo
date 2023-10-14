//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morninng/z-go-react-monorepo/internal/application"
	"github.com/morninng/z-go-react-monorepo/internal/infrastructure/datastore"
	"github.com/morninng/z-go-react-monorepo/internal/interfaces/handlers"
)

func InitializeTaskHandler() *handlers.TaskHandler {
	wire.Build(
		datastore.SuperSet,
		application.SuperSet,
		handlers.SuperSet,
	)
	return nil
}
