package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(authControl *controllers.AuthCotroller, produkcontrol *controllers.ProductIkanController, filterproduk *controllers.ProductFilterControl, detailproduk *controllers.ProductDetailControl, cartProduk *controllers.KeranjangControl) *gin.Engine {
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
		// cartRoutes.POST("/checkout", cartController.Checkout)
	}
	return r
}
