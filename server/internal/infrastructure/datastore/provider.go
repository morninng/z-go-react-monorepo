package datastore

import "github.com/google/wire"

var SuperSet = wire.NewSet(
	NewTaskRepository,
	NewDB,
)
