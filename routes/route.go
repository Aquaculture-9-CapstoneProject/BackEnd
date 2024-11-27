package routes

import (
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(authcontrol *controllers.AuthCotroller) *gin.Engine {
	r := gin.Default()
	r.GET("/")         //Home Utama
	r.GET("/register") //Route register render fe
	r.GET("/login")    // Route Login render fe
	r.POST("/register", authcontrol.DaftarAkun)
	r.POST("/login", authcontrol.Login)
	return r
}
