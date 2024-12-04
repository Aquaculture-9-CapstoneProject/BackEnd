package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routes(authControl *controllers.AuthCotroller, produkcontrol *controllers.ProductIkanController, artikelControl *controllers.ArtikelController) *gin.Engine {
	r := gin.Default()

	// Tambahkan middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://blue-bay.vercel.app"},
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

	r.GET("/produk-termurah", produkcontrol.GetTermurahProduk)
	r.GET("/produk-populer", produkcontrol.GetPopulerProduk)

	r.GET("/artikel", artikelControl.GetAll)
	r.GET("/artikel/:id", artikelControl.GetDetails)
	r.POST("/artikel", artikelControl.Create)
	r.PUT("/artikel/:id", artikelControl.Update)
	r.DELETE("/artikel/:id", artikelControl.Delete)

	return r
}
