package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(authControl *controllers.AuthCotroller) *gin.Engine {
	r := gin.Default()

	r.POST("/register", authControl.DaftarAkun)
	r.POST("/login", authControl.Login)

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
