package main

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/config"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	admincontroller "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers/adminController"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories/admin"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/routes"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	adminservices "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services/adminServices"
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
	adminDasbordRepo := admin.NewAdminPaymentRepository(db)

	authServis := services.NewAuthUseCase(authRepo)
	filterService := services.NewProductFilterService(filterRepo)
	OrderDetailServices := services.NeworderService(OrderDetailRepo, detailProdukRepo)
	productService := services.NewProductIkanServices(productRepo)
	CartServices := services.NewServicesKeranjang(CartRepo, detailProdukRepo, OrderDetailRepo)
	PaymentServices := services.NewPaymentServices(PaymentRepo)
	detailProdukServices := services.NewProductDetailServices(detailProdukRepo, ReviewRepo)
	ReviewServices := services.NewServiceRating(ReviewRepo, detailProdukServices)
	adminDasbordServices := adminservices.NewAdminPaymentService(adminDasbordRepo)

	authControl := controllers.NewAuthController(authServis)
	productController := controllers.NewProductIkanController(productService)
	filterController := controllers.NewProductFilterControl(filterService)
	OrderDetailController := controllers.NewOrderControl(OrderDetailServices)
	CartController := controllers.NewCartControl(CartServices)
	PaymentController := controllers.NewPaymentController(PaymentServices)
	ReviewController := controllers.NewReviewController(ReviewServices)
	detailProdukControl := controllers.NewProductDetailControl(detailProdukServices)
	adminDasboardController := admincontroller.NewAdminPaymentController(adminDasbordServices)

	r := routes.Routes(authControl, productController, filterController, detailProdukControl, CartController, OrderDetailController, PaymentController, ReviewController, adminDasboardController)

	r.Run(":8000")
}
