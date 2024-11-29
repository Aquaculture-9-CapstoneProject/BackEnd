package controllers

import (
	"net/http"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/middlewares"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/services"
	"github.com/gin-gonic/gin"
)

type AuthCotroller struct {
	usecase services.AuthUseCase
}

func NewAuthController(usecase services.AuthUseCase) *AuthCotroller {
	return &AuthCotroller{usecase: usecase}
}

func (ctrl *AuthCotroller) DaftarAkun(c *gin.Context) {
	var input struct {
		NamaLengkap string `json:"namalengkap"`
		Alamat      string `json:"alamat"`
		NoTelpon    string `json:"notelpon"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		KonfirPass  string `json:"konfirpass"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input gagal", "status": false})
		return
	}
	if len(input.Password) <= 5 {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Password Minimal 6 Karakter"})
		return
	}
	user, err := ctrl.usecase.DaftarUser(input.NamaLengkap, input.Alamat, input.NoTelpon, input.Email, input.Password, input.KonfirPass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": user.Email + " Berhasil Didaftarkan", "status": true})
}

func (ctrl *AuthCotroller) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input gagal", "status": false})
		return
	}

	user, err := ctrl.usecase.LoginUser(input.Email, input.Password)
	if err == nil && user != nil {
		tokenUser, err := middlewares.GenerateJwt(user.ID, "user")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Berhasil Login",
			"status":  "User",
			"token":   tokenUser,
			"user": gin.H{
				"email":    user.Email,
				"nama":     user.NamaLengkap,
				"alamat":   user.Alamat,
				"noTelpon": user.NoTelpon,
			},
		})
		return

	}

	admin, err := ctrl.usecase.LoginAdmin(input.Email, input.Password)
	if err == nil && admin != nil {
		tokenAdmin, err := middlewares.GenerateJwt(admin.ID, "admin")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Email": admin.Email, "Message": "Berhasil Login", "Status": "Admin", "Token": tokenAdmin})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"Message": "Email atau password salah", "status": false})
}

func (ctrl *AuthCotroller) Logout(c *gin.Context) {
	// Menghapus token dari client-side (frontend)
	// di frontend token dihapus dari local storage atau session storage
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout berhasil",
		"status":  true,
	})
}
