package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/krissukoco/go-project-layout/internal/config"
	auth_handler "github.com/krissukoco/go-project-layout/internal/delivery/http/handler/auth"
	profile_handler "github.com/krissukoco/go-project-layout/internal/delivery/http/handler/profile"
	"github.com/krissukoco/go-project-layout/internal/delivery/http/middleware"
	http_response "github.com/krissukoco/go-project-layout/internal/delivery/http/response"
	user_repository_impl_pg "github.com/krissukoco/go-project-layout/internal/repository/user/impl_pg"
	auth_token_usecase_impl "github.com/krissukoco/go-project-layout/internal/usecase/auth_token/impl"
	auth_user_usecase_impl "github.com/krissukoco/go-project-layout/internal/usecase/auth_user/impl"
	user_usecase_impl "github.com/krissukoco/go-project-layout/internal/usecase/user/impl"
	_ "github.com/lib/pq"
)

func connectPostgres(host, username, password, dbname string, port uint, enableSsl bool, timezone string) (*sql.DB, error) {
	ssl := "disable"
	if enableSsl {
		ssl = "enable"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s timezone=%s", host, username, password, dbname, port, ssl, timezone)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Database
	db, err := connectPostgres(
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDbname,
		cfg.PostgresPort,
		cfg.PostgresEnableSsl,
		cfg.PostgresTimezone,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	userRepo := user_repository_impl_pg.New(db)

	// Usecase
	authTokenUc := auth_token_usecase_impl.New(cfg.JwtSecret, userRepo)
	authUserUc := auth_user_usecase_impl.New(userRepo)
	userUc := user_usecase_impl.New(userRepo)

	// Handlers
	authHandler := auth_handler.New(authUserUc)
	profileHandler := profile_handler.New(userUc)

	// Middlewares
	authMw := middleware.NewAuthMiddleware(authTokenUc)

	// Server
	app := fiber.New(fiber.Config{
		ErrorHandler: http_response.NewErrorHandler(true),
	})
	app.Use(recover.New())

	v1 := app.Group("/v1")
	v1.Use(logger.New())
	v1.Use(cors.New())
	{
		// Auth routes
		r := v1.Group("/auth")
		r.Post("/login", authHandler.Login)
	}
	{
		// Profile routes
		r := v1.Group("/profile")
		r.Use(authMw)
		r.Get("/me", profileHandler.GetProfile)
	}

	// Gracefully start server
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Server %s started on port %d ...\n", cfg.ServiceName, cfg.Port)
	<-exit

	log.Println("Shutdown signal received. Shutting down...")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		app.Shutdown()
		log.Println("Fiber App shut down")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		db.Close()
		log.Println("PostgreSQL database connection closed")
		wg.Done()
	}()

	// Terminate if too long in shutting down server
	go func() {
		timeout := time.Duration(cfg.GracefulTimeout) * time.Second
		time.Sleep(timeout)
		log.Fatalf("Termination process took longer than %v. Force exit!", timeout)
	}()

	wg.Wait()
	log.Println("Server gracefully shutdown")
}
