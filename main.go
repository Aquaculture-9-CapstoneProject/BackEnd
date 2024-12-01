package main

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/config"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/routes"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
)

// *Branch Main
// *Dokumentasi Postman
// *https://www.bluebay.my.id/route

func main() {
	db := config.CreateDatabase()
	authRepo := repositories.NewAuthRepo(db)
	authServis := services.NewAuthUseCase(authRepo)
	authControl := controllers.NewAuthController(authServis)

	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductUseCase(productRepo)
	productController := controllers.NewProductController(productService)

	ratingRepo := repositories.NewRatingRepo(db)
	ratingService := services.NewRatingUseCase(ratingRepo)
	ratingController := controllers.NewRatingController(ratingService)

	r := routes.Routes(authControl, productController, ratingController)

	r.Run(":8000")
}
