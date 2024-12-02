package database

type Driver string

const (
	SQLite   Driver = "sqlite"
	Postgres Driver = "postgres"
	Redis    Driver = "redis"
)

type Config struct {
	Driver   Driver
	Database string
	URL      string
	// 1. define database config
}

//database url patterns:
//url: "sqlite://schedules.db"
//# or
//url: "postgres://user:pass@localhost:5432/schedules"
//# or
//url: "redis://:pass@localhost:6379/0"
