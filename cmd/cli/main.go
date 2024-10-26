package main

import (
	"github.com/admiralyeoj/animanager/internal/config"
	"github.com/admiralyeoj/animanager/internal/database"
	"github.com/admiralyeoj/animanager/internal/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.CloseLogger()

	// load .env file
	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file", zap.Error(err))
	}

	cfg := config.NewConfig()

	err = cfg.ParseFlags()
	if err != nil {
		logger.Logger.Fatal("Failed to parse command-line flags", zap.Error(err))
	}

	db, err := database.Connect(cfg)
	if err != nil {
		logger.Logger.Fatal("Failed to connect to the database", zap.Error(err))
		panic(err)
	}
	defer database.Close(db)

	// fmt.Println("Connected to database")

	startRepl(cfg, db)
}
