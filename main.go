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

	filterRepo := repositories.NewProductFilterRepo(db)
	filterService := services.NewProductFilterService(filterRepo)
	filterController := controllers.NewProductFilterControl(filterService)

	detailProdukRepo := repositories.NewProductDetailRepo(db)
	detailProdukServices := services.NewProductDetailServices(detailProdukRepo)
	detailProdukControl := controllers.NewProductDetailControl(detailProdukServices)

	CartRepo := repositories.NewKeranjangRepo(db)
	CartServices := services.NewServicesKeranjang(CartRepo, detailProdukRepo)
	CartController := controllers.NewCartControl(CartServices)

	OrderDetailRepo := repositories.NewOrderRepo(db)
	OrderDetailServices := services.NeworderService(OrderDetailRepo, detailProdukRepo)
	OrderDetailController := controllers.NewOrderControl(OrderDetailServices)

	r := routes.Routes(authControl, productController, filterController, detailProdukControl, CartController, OrderDetailController)

	r.Run(":8000")
}
