package controllers

import (
	"fmt"
	"labpro/single-service/initializers"
	"labpro/single-service/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSelf(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		CreateResponse(c, http.StatusUnauthorized, "error", "Authorization header is missing.", nil)
		return
	}

	claims := ParseJWT(token)

	if !IsClaimsValid(claims) {
		CreateResponse(c, http.StatusUnauthorized, "error", "Invalid or expired token.", nil)
		return
	}

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		CreateResponse(c, http.StatusUnauthorized, "error", "Invalid or expired token.", nil)
		return
	}

	var user models.User
	fmt.Println(claims["user_id"])
	res := initializers.DB.First(&user, "id = ?", claims["user_id"])

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error == gorm.ErrRecordNotFound {
			CreateResponse(c, http.StatusOK, "success", "There is no user with certain ID.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve user.", nil)
		}
		return
	}

	self := gin.H{
		"username": user.Username,
		"name":     user.Nama,
	}
	CreateResponse(c, http.StatusOK, "success", "Self retrieved successfully.", self)
}

func GetBarang(c *gin.Context) {
	q := c.Query("q")
	perusahaan := c.Query("perusahaan")

	var arrBarang []models.Barang
	res := initializers.DB.Where("nama LIKE ? OR kode LIKE ?", "%"+q+"%", "%"+q+"%")
	if perusahaan != "" {
		res = res.Where("perusahaan_id = ?", perusahaan)
	}
	res = res.Find(&arrBarang)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve barang.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Barang retrieved successfully.", arrBarang)
}

func GetBarangById(c *gin.Context) {
	id := c.Param("id")

	var barang models.Barang
	res := initializers.DB.First(&barang, "id = ?", id)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error == gorm.ErrRecordNotFound {
			CreateResponse(c, http.StatusNotFound, "success", "There is no barang with certain ID.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve barang.", nil)
		}
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Barang retrieved successfully.", gin.H{
		"id":            barang.ID,
		"nama":          barang.Nama,
		"harga":         barang.Harga,
		"stok":          barang.Stok,
		"kode":          barang.Kode,
		"perusahaan_id": barang.PerusahaanID,
	})
}

func GetPerusahaan(c *gin.Context) {
	q := c.Query("q")

	var arrPerusahaan []models.Perusahaan
	res := initializers.DB.Where("nama LIKE ? OR kode LIKE ?", "%"+q+"%", "%"+q+"%").Find(&arrPerusahaan)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve perusahaan.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Perusahaan retrieved successfully.", arrPerusahaan)
}

func GetPerusahaanById(c *gin.Context) {
	id := c.Param("id")

	var perusahaan models.Perusahaan
	res := initializers.DB.First(&perusahaan, "id = ?", id)

	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error == gorm.ErrRecordNotFound {
			CreateResponse(c, http.StatusNotFound, "success", "There is no perusahaan with certain ID.", nil)
			// CreateResponse(c, http.StatusOK, "success", "There is no perusahaan with certain ID.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve perusahaan.", nil)
		}
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Perusahaan retrieved successfully.", perusahaan)
}
