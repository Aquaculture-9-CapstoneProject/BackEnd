package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	admincontroller "github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers/adminController"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(authControl *controllers.AuthCotroller, produkcontrol *controllers.ProductIkanController, filterproduk *controllers.ProductFilterControl, detailproduk *controllers.ProductDetailControl, cartProduk *controllers.KeranjangControl, orderProduk *controllers.OrderControl, payment *controllers.PaymentControl, review *controllers.ReviewController, dasboard *admincontroller.AdminPaymentController) *gin.Engine {
	r := gin.Default()

	// Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://blue-bay.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rute untuk register dan login
	r.POST("/register", authControl.DaftarAkun)
	r.POST("/login", authControl.Login)
	r.POST("/logout", authControl.Logout)

	r.GET("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halaman Register untuk Frontend",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halaman Login untuk Frontend",
		})
	})
	route := r.Group("/")
	route.Use(middlewares.JWTAuth())
	route.GET("/products", filterproduk.FilterProduct)

	route.GET("/produk-termurah", produkcontrol.GetTermurahProduk)
	route.GET("/produk-populer", produkcontrol.GetPopulerProduk)
	route.GET("/produk", produkcontrol.GetProductAll)

	route.GET("/products/:id", detailproduk.CekDetailProdukByID)

	route.POST("/products/:product_id/reviews", review.AddReview)
	route.GET("/products/:id/reviews", review.GetReviewsByProduct)

	cartRoutes := route.Group("/cart")
	{
		cartRoutes.POST("/tambah", cartProduk.AddToCart)
		cartRoutes.GET("", cartProduk.GetCartUser)
		cartRoutes.DELETE("/:cartID", cartProduk.DeleteKeranjang)
		cartRoutes.POST("/checkout", cartProduk.CheckOut)
	}

	orderRoutes := route.Group("/orders")
	{
		orderRoutes.POST("", orderProduk.PlaceOrder)
		orderRoutes.GET("/checkout", orderProduk.GetOrderForCheckout)
	}

	paymentRoutes := route.Group("/payments")
	{
		// Endpoint untuk membuat pembayaran
		paymentRoutes.POST("", payment.TambahPayment)

		// Endpoint untuk mengecek status pembayaran
		paymentRoutes.GET("/:invoiceID/status", payment.CheckPaymentStatus)

		// Endpoint untuk membatalkan pembayaran
		paymentRoutes.POST("/cancel", payment.CancelPayment)

		// Endpoint untuk mendapatkan detail pembayaran berdasarkan Invoice ID
		paymentRoutes.GET("/detail/:invoiceID", payment.GetPaymentByInvoiceID)

		// Endpoint untuk mendapatkan pesanan yang sudah dibayar
		paymentRoutes.GET("/order/paid", payment.GetPaidOrders)
	}

	adminRoute := route.Group("/admin", middlewares.AdminOnly())
	adminRoute.GET("/totalpendapatan", dasboard.GetAdminTotalPendapatanBulanIni)
	adminRoute.GET("/totalpesanan", dasboard.GetAdminJumlahPesananBulanIni)
	adminRoute.GET("/totalproduk", dasboard.GetTotalProduk)
	adminRoute.GET("/statustransaksi", dasboard.GetJumlahStatus)
	adminRoute.GET("/totaldikirim", dasboard.GetJumlahPesananDikirim)
	adminRoute.GET("/totalditerima", dasboard.GetJumlahPesananDiterima)
	adminRoute.GET("/totalpendapatan", dasboard.TampilkanTotalPendapatan)

	return r
}
