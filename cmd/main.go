package main

import (
	"log"

	"project/internal/domain/models"
	"project/internal/infrastructure/db"
	"project/internal/infrastructure/repository"
	"project/internal/interfaces/http"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// Connect to the database
	dbConn := db.ConnectDatabase()

	// Automatically migrate your schema, keeping the database in sync with your models
	runMigrations(dbConn)

	// Initialize Repositories
	userRepo := repository.NewUserRepository(dbConn)
	groupRepo := repository.NewGroupRepository(dbConn)
	roleRepo := repository.NewRoleRepository(dbConn)

	// Initialize Usecases
	userUsecase := usecase.NewUserUsecase(userRepo, groupRepo)
	groupUsecase := usecase.NewGroupUsecase(groupRepo)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)

	// Initialize Handlers
	http.NewUserHandler(r, userUsecase)
	http.NewGroupHandler(r, groupUsecase)
	http.NewRoleHandler(r, roleUsecase)

	// Run the Gin server
	r.Run(":8080")
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Role{},
		&models.UserGroup{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
