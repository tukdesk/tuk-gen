package db

const (
	MYSQL      = "mysql"
	POSTGRESQL = "postgres"
	SQLITE3    = "sqlite3"
)

var defaultEngine = MYSQL

func DefaultEngine() string {
	return defaultEngine
}

func SetDefaultEngine(engine string) {
	defaultEngine = engine
}
