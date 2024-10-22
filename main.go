package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/admiralyeoj/anime-announcements/configs"
	"github.com/admiralyeoj/anime-announcements/internal/aniListApi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// load .env file
	godotenv.Load()

	postgresURI := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		log.Panic(err)
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Panic(err)
	}

	fmt.Println("Connected to database")

	aniListClient := aniListApi.NewClient(time.Second * 5)

	cfg := &configs.AniListConfig{
		AniListClient: aniListClient,
	}

	startRepl(cfg)
}
