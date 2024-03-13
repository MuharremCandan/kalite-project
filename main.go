package main

import (
	"go-backend-test/pkg/api/router"
	"go-backend-test/pkg/config"
	database "go-backend-test/pkg/db"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	_, err = database.NewPgDb(config).ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	gin.ForceConsoleColor()

	r := gin.Default()

	router.NewRouter().SetUpRouter(r)

	s := &http.Server{
		Addr:           net.JoinHostPort(config.HttpServer.Host, config.HttpServer.Port),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("server has started at : ", net.JoinHostPort(config.HttpServer.Host, config.HttpServer.Port))
	log.Fatal("create server error: ", s.ListenAndServe())

}
