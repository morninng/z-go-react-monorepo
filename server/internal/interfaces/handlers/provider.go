package handlers

import "github.com/google/wire"

var SuperSet = wire.NewSet(
	NewTaskHandler,
)
