package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(authControl *controllers.AuthCotroller, produkcontrol *controllers.ProductIkanController, filterproduk *controllers.ProductFilterControl, detailproduk *controllers.ProductDetailControl, cartProduk *controllers.KeranjangControl, orderProduk *controllers.OrderControl, chatControl *controllers.ChatController, artikelControl *controllers.ArtikelController, adminProductControl *controllers.AdminProductController) *gin.Engine {
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

	chatRoutes := route.Group("/chats")
	{
		chatRoutes.GET("", chatControl.GetAllChats)
		chatRoutes.POST("", chatControl.ChatController)
	}

	r.GET("/artikel", artikelControl.GetAll)
	r.GET("/artikel/:id", artikelControl.GetDetails)
	r.POST("/artikel", artikelControl.Create)
	r.PUT("/artikel/:id", artikelControl.Update)
	r.DELETE("/artikel/:id", artikelControl.Delete)

	r.GET("/dashboard/products", adminProductControl.SearchAdminProducts)
	r.GET("/dashboard/products/:id", adminProductControl.GetAdminProductDetails)
	r.POST("/dashboard/products", adminProductControl.CreateAdminProduct)
	r.PUT("/dashboard/products/:id", adminProductControl.UpdateAdminProduct)
	r.DELETE("/dashboard/products/:id", adminProductControl.DeleteAdminProduct)

	return r
}
