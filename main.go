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

	productRepo := repositories.NewProdukIkanRepo(db)
	productService := services.NewProductIkanServices(productRepo)
	productController := controllers.NewProductIkanController(productService)

	artikelRepo := repositories.NewArtikelRepo(db)
	artikelService := services.NewArtikelService(artikelRepo)
	artikelController := controllers.NewArtikelController(artikelService)

	r := routes.Routes(authControl, productController, artikelController)

	r.Run(":8000")
}
