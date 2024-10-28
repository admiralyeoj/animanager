package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/admiralyeoj/animanager/cmd/cli"
	"github.com/admiralyeoj/animanager/cmd/cron"
	"github.com/admiralyeoj/animanager/internal/config"
	"github.com/admiralyeoj/animanager/internal/database"
	"github.com/admiralyeoj/animanager/internal/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// Define a flag to set the mode (cli or cron)
	mode := flag.String("mode", "cli", "Application mode: 'cli' for command line or 'cron' for cron job")
	flag.Parse()

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

	switch *mode {
	case "cli":
		cli.StartCli(cfg, db)
	case "cron":
		cron.StartCron(cfg, db)
	default:
		fmt.Println("Invalid mode. Please choose 'cli' or 'cron'")
		os.Exit(1)
	}

}
