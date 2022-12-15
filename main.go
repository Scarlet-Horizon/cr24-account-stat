package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"main/controller"
	"main/db"
	_ "main/docs"
	"main/env"
	"main/util"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//	@title			cr24 Account Statistics API
//	@version		1.0
//	@description	API for account statistics for cr24-account-service project
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	David Slatinek
//	@contact.url	https://github.com/david-slatinek

//	@accept		json
//	@produce	json
//	@schemes	http

//	@license.name	GNU General Public License v3.0
//	@license.url	https://www.gnu.org/licenses/gpl-3.0.html

//	@securityDefinitions.apikey	JWT
//@in header
//@name Authorization

//	@host		localhost:8090
//	@BasePath	/api/v1
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

	accountController := controller.AccountController{
		DB: &db.AccountDB{
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
		api.GET("/all", statController.All)

		api.POST("/account", accountController.Create)
		api.GET("/account", accountController.Get)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
