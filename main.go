package main

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/config"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/routes"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-contrib/cors"
)

func main() {
	db := config.CreateDatabase()
	authRepo := repositories.NewAuthRepo(db)
	authServis := services.NewAuthUseCase(authRepo)
	authControl := controllers.NewAuthController(authServis)

	r := routes.Routes(authControl)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Run(":8000")
}
