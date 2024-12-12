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

	chatRepo := repositories.NewChatRepo(db)
	chatService := services.NewChatService(chatRepo)
	chatController := controllers.NewChatController(chatService)

	artikelRepo := repositories.NewArtikelRepo(db)
	artikelService := services.NewArtikelService(artikelRepo)
	artikelController := controllers.NewArtikelController(artikelService)

	adminProductRepo := repositories.NewAdminProductRepo(db)
	adminProductService := services.NewAdminProductService(adminProductRepo)
	adminProductController := controllers.NewAdminProductController(adminProductService)

	r := routes.Routes(authControl, productController, filterController, detailProdukControl, CartController, OrderDetailController, chatController, artikelController, adminProductController)

	r.Run(":8000")
}
