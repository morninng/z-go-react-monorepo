//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/morninng/z-go-react-monorepo/internal/application"
	"github.com/morninng/z-go-react-monorepo/internal/handlers"
	"github.com/morninng/z-go-react-monorepo/internal/infrastructure/datastore"
)

func InitializeTaskHandler() *handlers.TaskHandler {
	wire.Build(
		datastore.SuperSet,
		application.SuperSet,
		handlers.SuperSet,
	)
	return nil
}
