package database

type Driver interface {
}
type Config struct {
	Driver Driver
	URL    string
	// 1. define database config
}

//database url patterns:
//url: "sqlite://schedules.db"
//# or
//url: "postgres://user:pass@localhost:5432/schedules"
//# or
//url: "redis://:pass@localhost:6379/0"
