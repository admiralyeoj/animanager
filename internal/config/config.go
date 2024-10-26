package config

import (
	"flag"
	"os"

	aniListRepo "github.com/admiralyeoj/anime-announcements/internal/aniList/repository"
	dbRepo "github.com/admiralyeoj/anime-announcements/internal/database/repository"
)

// Config holds all the configuration settings for the application.
type Config struct {
	Name    string
	Port    string
	Env     string
	DB      DBConfig
	Limiter LimiterConfig
	SMTP    SMTPConfig
}

func NewConfig() *Config {
	return &Config{}
}

type Repositories struct {
	AniListRepo   aniListRepo.AniListRepository
	DatabaseRepos *dbRepo.DatabaseRepositories
	// Add other repositories here
}

// ParseFlags parses the environment variables and flags into the Config struct.
func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Name, "name", os.Getenv("APP_NAME"), "App Name")
	flag.StringVar(&cfg.Port, "port", os.Getenv("APP_PORT"), "API server port")
	flag.StringVar(&cfg.Env, "env", os.Getenv("APP_ENV"), "Environment (local|development|staging|production)")

	// Database configuration
	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")
	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	// Limiter configuration
	flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	// SMTP configuration
	flag.StringVar(&cfg.SMTP.Host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.IntVar(&cfg.SMTP.Port, "smtp-port", 25, os.Getenv("SMTP_PORT"))
	flag.StringVar(&cfg.SMTP.Username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "openMovie <no-reply@test.user.net>", "SMTP sender")

	flag.Parse()
	return nil
}
