package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/zaenalarifin12/test-ecommerce/internal/api"
	"github.com/zaenalarifin12/test-ecommerce/internal/config"
	"github.com/zaenalarifin12/test-ecommerce/internal/repository"
	"github.com/zaenalarifin12/test-ecommerce/internal/service"
	"github.com/zaenalarifin12/test-ecommerce/internal/utils"
	"log"
)

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		panic("can't load config")
	}

	runMigrationDB(conf.MigrationUrl, conf.DBSource)
	pool, err := connectToDatabase(conf)

	userRepository := repository.NewUser(pool)
	userService := service.NewUser(userRepository)

	productRepo := repository.NewProductRepository(pool)
	productService := service.NewProduct(productRepo)

	// setup gin
	router := gin.Default()

	// list router
	api.NewUser(router, userService)
	//api.NewUser(router, userService, middleware.JWTMiddleware())

	// Register middleware for JWT authentication
	authMiddleware := utils.JWTMiddleware(conf.SecretKey)

	// Register product API routes
	api.NewProduct(router, productService, authMiddleware)

	port := fmt.Sprintf(":%v", conf.ServerPort)

	err = router.Run(port)
	if err != nil {
		panic("can't start server")
	}
}

func connectToDatabase(conf config.Config) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(conf.DBSource)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	return pool, nil
}

func runMigrationDB(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance: ", err)
	}

	UpMigration(migration)
	//rollbackMigration(migration)
}

func rollbackMigration(migration *migrate.Migrate) {
	// Perform rollback
	if err := migration.Down(); err != nil {
		log.Fatal("failed to rollback migration: ", err)
	}
	log.Println("rollback successful")
}

func UpMigration(migration *migrate.Migrate) {
	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migrate up! ", err)
	}

	log.Println("db migrate successfully")

}
