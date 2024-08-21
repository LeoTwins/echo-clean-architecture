package main

import (
	"log"

	"github.com/LeoTwins/go-clean-architecture/cmd/api/router"
	"github.com/LeoTwins/go-clean-architecture/internal/infrastructure/db"
	"github.com/LeoTwins/go-clean-architecture/pkg/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
	}

	db, err := db.NewDB(cfg.DBConfig)
	if err != nil {
		log.Fatalf("データベースとの接続に失敗しました: %v", err)
	}

	e := echo.New()
	router.NewRouter(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
