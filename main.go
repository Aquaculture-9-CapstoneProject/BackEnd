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
	adminTransaksiRepo := admin.NewAdminTransaksiRepo(db)
	adminPesananRepo := admin.NewPesananRepo(db)
	adminFilterRepo := admin.NewAdminFilterRepo(db)
	profileRepo := repositories.NewProfileRepository(db)

	authServis := services.NewAuthUseCase(authRepo)
	filterService := services.NewProductFilterService(filterRepo)
	OrderDetailServices := services.NeworderService(OrderDetailRepo, detailProdukRepo)
	productService := services.NewProductIkanServices(productRepo)
	CartServices := services.NewServicesKeranjang(CartRepo, detailProdukRepo, OrderDetailRepo)
	PaymentServices := services.NewPaymentServices(PaymentRepo)
	detailProdukServices := services.NewProductDetailServices(detailProdukRepo, ReviewRepo)
	ReviewServices := services.NewServiceRating(ReviewRepo, detailProdukServices)
	adminDasbordServices := adminservices.NewAdminPaymentService(adminDasbordRepo)
	adminTransaksiServices := adminservices.NewAdminTransaksiServices(adminTransaksiRepo)
	adminPesananServices := adminservices.NewPesananServices(adminPesananRepo)
	adminFilterServices := adminservices.NewAdminFilterServices(adminFilterRepo)
	profileServices := services.NewProfileService(profileRepo)
	ExportServices := adminservices.NewProducExportService(productRepo)

	authControl := controllers.NewAuthController(authServis)
	productController := controllers.NewProductIkanController(productService)
	filterController := controllers.NewProductFilterControl(filterService)
	OrderDetailController := controllers.NewOrderControl(OrderDetailServices)
	CartController := controllers.NewCartControl(CartServices)
	PaymentController := controllers.NewPaymentController(PaymentServices)
	ReviewController := controllers.NewReviewController(ReviewServices)
	detailProdukControl := controllers.NewProductDetailControl(detailProdukServices)
	adminDasboardController := admincontroller.NewAdminPaymentController(adminDasbordServices)
	adminTransakasiController := admincontroller.NewAdminTransaksiController(adminTransaksiServices)
	adminPesananController := admincontroller.NewAdminPesananController(adminPesananServices)
	adminFilterController := admincontroller.NewAdminFilterController(adminFilterServices)
	profileController := controllers.NewProfileController(profileServices)
	ExportController := admincontroller.NewProductExportController(ExportServices)

	chatRepo := repositories.NewChatRepo(db)
	chatService := services.NewChatService(chatRepo)
	chatController := controllers.NewChatController(chatService)

	artikelRepo := repositories.NewArtikelRepo(db)
	artikelService := services.NewArtikelService(artikelRepo)
	artikelController := controllers.NewArtikelController(artikelService)

	adminProductRepo := admin.NewAdminProductRepo(db)
	adminProductService := adminservices.NewAdminProductService(adminProductRepo)
	adminProductController := admincontroller.NewAdminProductController(adminProductService)

	adminArtikelRepo := admin.NewAdminArtikelRepo(db)
	adminArtikelService := adminservices.NewAdminArtikelService(adminArtikelRepo)
	adminArtikelController := admincontroller.NewAdminArtikelController(adminArtikelService)

	r := routes.Routes(authControl, productController, filterController, detailProdukControl, CartController, OrderDetailController, PaymentController, ReviewController, adminDasboardController, artikelController, chatController, adminProductController, adminTransakasiController, adminPesananController, adminFilterController, profileController, ExportController, adminArtikelController)

	r.Run(":8000")
}

// BESOK CRUD PRODUK DAN UNIT TES
// export csv pada produk done
