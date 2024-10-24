package config

// DBConfig holds the database-related configuration settings.
type DBConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
