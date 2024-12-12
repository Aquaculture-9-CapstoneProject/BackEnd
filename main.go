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
//

func main() {
	db := config.CreateDatabase()
	authRepo := repositories.NewAuthRepo(db)
	detailProdukRepo := repositories.NewProductDetailRepo(db)
	ReviewRepo := repositories.NewRatingRepo(db)
	PaymentRepo := repositories.NewPaymentRepo(db)
	CartRepo := repositories.NewKeranjangRepo(db)
	OrderDetailRepo := repositories.NewOrderRepo(db)
	filterRepo := repositories.NewProductFilterRepo(db)
	productRepo := repositories.NewProdukIkanRepo(db)

	authServis := services.NewAuthUseCase(authRepo)
	filterService := services.NewProductFilterService(filterRepo)
	OrderDetailServices := services.NeworderService(OrderDetailRepo, detailProdukRepo)
	productService := services.NewProductIkanServices(productRepo)
	CartServices := services.NewServicesKeranjang(CartRepo, detailProdukRepo, OrderDetailRepo)
	PaymentServices := services.NewPaymentServices(PaymentRepo)
	detailProdukServices := services.NewProductDetailServices(detailProdukRepo, ReviewRepo)
	ReviewServices := services.NewServiceRating(ReviewRepo, detailProdukServices)

	authControl := controllers.NewAuthController(authServis)
	productController := controllers.NewProductIkanController(productService)
	filterController := controllers.NewProductFilterControl(filterService)
	OrderDetailController := controllers.NewOrderControl(OrderDetailServices)
	CartController := controllers.NewCartControl(CartServices)
	PaymentController := controllers.NewPaymentController(PaymentServices)
	ReviewController := controllers.NewReviewController(ReviewServices)
	detailProdukControl := controllers.NewProductDetailControl(detailProdukServices)

	chatRepo := repositories.NewChatRepo(db)
	chatService := services.NewChatService(chatRepo)
	chatController := controllers.NewChatController(chatService)

	artikelRepo := repositories.NewArtikelRepo(db)
	artikelService := services.NewArtikelService(artikelRepo)
	artikelController := controllers.NewArtikelController(artikelService)

	r := routes.Routes(authControl, productController, filterController, detailProdukControl, CartController, OrderDetailController, PaymentController, ReviewController, chatController, artikelController)

	r.Run(":8000")
}
