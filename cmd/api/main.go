package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"

	"cap-api/internal/handler"
	"cap-api/internal/service"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./memos.db"
	}

	db, err := sqlx.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schemaBytes, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(string(schemaBytes)); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	memoSvc := service.NewMemoService(db)
	memoHandler := handler.NewMemoHandler(memoSvc)
	memoHandler.Register(r)

	log.Println("server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
