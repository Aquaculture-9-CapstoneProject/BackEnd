package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/gin-gonic/gin"
)

// Menyusun rute API untuk AuthController
func Routes(authControl *controllers.AuthCotroller) *gin.Engine {
	r := gin.Default()

	// Menentukan rute API untuk register dan login
	r.POST("/register", authControl.DaftarAkun) // POST untuk register
	r.POST("/login", authControl.Login)         // POST untuk login

	// Jika Anda ingin memberikan halaman atau template frontend untuk register dan login, gunakan GET
	// Namun, untuk API, ini tidak perlu. Hapus jika hanya berfungsi sebagai API.
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

	return r
}
