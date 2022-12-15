package main

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/env"
	"os"
	"time"
)

func main() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatalf("failed to load env variables: %v", err)
	}

	uri := os.Getenv("MYSQL_URL")
	if uri == "" {
		log.Fatal("MYSQL_URL is not set")
	}

	mysqlDB, err := sql.Open("mysql", os.Getenv("MYSQL_URL"))
	if err != nil {
		log.Fatalf("error with sql.Open: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = mysqlDB.PingContext(ctx)
	if err != nil {
		log.Fatalf("ping error: %v", err)
	}

	mysqlDB.SetConnMaxLifetime(time.Minute * 3)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(10)

	defer func(mysqlDB *sql.DB) {
		if err := mysqlDB.Close(); err != nil {
			log.Printf("Close() error: %v", err)
		}
	}(mysqlDB)

	gin.SetMode(os.Getenv("GIN_MODE"))
}
