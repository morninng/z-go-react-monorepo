package application

import "github.com/google/wire"

var SuperSet = wire.NewSet(
	NewTaskApplicationService,
)
