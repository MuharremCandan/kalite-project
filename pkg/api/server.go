package api

import (
	"context"
	"fmt"
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
	"gorm.io/gorm"
)

type Server struct {
	db         *gorm.DB
	config     *config.Config
	tokenMaker *token.Maker
	engine     *gin.Engine
}

func NewServer(config *config.Config) (*Server, error) {
	//genereta token maker
	maker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("could not create maker: %w", err)
	}
	//Connection of Postgresql
	pgDB := database.NewPgDb(config)
	db, err := pgDB.ConnectDb()
	if err != nil {
		return nil, fmt.Errorf("could not connect database: %w", err)
	}
	log.Println("Database Connection created successfully!")

	server := &Server{
		db:         db,
		config:     config,
		tokenMaker: &maker,
		engine:     gin.New(),
	}
	// container := dig.New()

	// container.Provide(repository.NewUserRepository(db))
	// container.Provide(service.NewUserService(container))
	// container.Provide(handler.NewUserHandler(container, maker, config))
	// container.Provide(repository.NewBrandRepository(db))
	// container.Provide(service.NewBrandService(container))
	// container.Provide(handler.NewBrandHandler(container))
	// container.Provide(handler.NewProductHandler(container))

	// if err := container.Invoke(container); err != nil {
	// 	return nil, fmt.Errorf("could not invoke container: %w", err)
	// }

	gin.ForceConsoleColor()

	//TODO: dependency injection container var ona bi bak
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, maker, config)

	productRepository := repository.NewProductRepository(db)
	prductService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(prductService)

	brandRepository := repository.NewBrandRepository(db)
	brandService := service.NewBrandService(brandRepository)
	brandHandler := handler.NewBrandHandler(brandService)

	router.NewRouter(userHandler, productHandler, brandHandler, config).SetUpRouter(server.engine)

	//server.engine.Use(middleware.AuthMiddleware(config.SecretKey))

	return server, nil
}

func (server *Server) StartServer(ctx context.Context) error {

	var err error
	ch := make(chan error, 1)
	go func() {

		s := &http.Server{
			Addr:           net.JoinHostPort(server.config.HttpServer.Host, server.config.HttpServer.Port),
			Handler:        server.engine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Println("server has started at : ", net.JoinHostPort(server.config.HttpServer.Host, server.config.HttpServer.Port))
		err := s.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to listendandserver. error: %w", err)
		}
		close(ch)
	}()
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		return net.ErrClosed
	}
}
