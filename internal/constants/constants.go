package constants

import (
	"time"
)

const (
	MaxDBPoolconns = 20
	MinDBPoolconns = 2

	PathOfMigrations = "internal/migrations"

	ShutdownTime = 10 * time.Second
)
