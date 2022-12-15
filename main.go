package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/controller"
	"main/db"
	"main/env"
	"main/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	statController := controller.StatController{
		DB: &db.StatDB{
			DB: mysqlDB,
		},
	}

	router := gin.Default()
	router.Use(util.CORS)
	api := router.Group("api/v1").Use(util.ValidateToken)
	{
		api.POST("/stat", statController.CreateStat)

		api.GET("/last", statController.LastEndpoint)
		api.GET("/most", statController.MostCalled)
	}

	srv := &http.Server{
		Addr:         ":8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		log.Println("server is up at: " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe() error: %s\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown() error: %v", err)
	}

	log.Println("shutting down")
}
