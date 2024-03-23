package main

import (
	"go-backend-test/pkg/api/handler"
	"go-backend-test/pkg/api/router"
	"go-backend-test/pkg/config"
	database "go-backend-test/pkg/db"
	"go-backend-test/pkg/repository"
	"go-backend-test/pkg/service"
	"go-backend-test/pkg/token"
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
	db, err := database.NewPgDb(config).ConnectDb()
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	gin.ForceConsoleColor()

	r := gin.Default()

	maker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		log.Fatalf("failed to create jwt maker: %v", err)
	}

	//TODO: dependency injection container var ona bi bak
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, maker, config)
	router.NewRouter(userHandler).SetUpRouter(r)

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
