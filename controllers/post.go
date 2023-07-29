package controllers

import (
	"fmt"
	"labpro/single-service/initializers"
	"labpro/single-service/models"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GenerateToken(userID string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = isAdmin
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Duration(initializers.Cfg.TOKEN_LIFESPAN_IN_HOURS) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(initializers.Cfg.SECRET_TOKEN))
}

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body.", nil)
		return
	}

	var user models.User
	res := initializers.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user)
	if res.Error != nil {
		fmt.Println(res.Error)
		if res.Error == gorm.ErrRecordNotFound {
			CreateResponse(c, http.StatusNotFound, "error", "There is no user with certain username and password.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to login.", nil)
		}
		return
	}

	token, err := GenerateToken(user.ID, user.Tipe == "admin")
	if err != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to generate token.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Token generated successfully.", gin.H{
		"user":  user,
		"token": token,
	})
}

func CreateBarang(c *gin.Context) {
	var req models.Barang

	if err := c.BindJSON(&req); err != nil {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body.", nil)
		return
	}

	if req.Harga == 0 {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body. Harga can't be 0.", nil)
		return
	}

	res := initializers.DB.Create(&req)
	if res.Error != nil {
		fmt.Println(res.Error)
		if strings.Contains(res.Error.Error(), "duplicate key") {
			CreateResponse(c, http.StatusConflict, "error", "Barang with the certain ID or kode is already exists.", nil)
		} else if strings.Contains(res.Error.Error(), "foreign key") {
			CreateResponse(c, http.StatusConflict, "error", "There is no perusahaan with the certain ID.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to create barang.", nil)
		}
		return
	}

	barang := gin.H{
		"id":            req.ID,
		"nama":          req.Nama,
		"harga":         req.Harga,
		"stok":          req.Stok,
		"kode":          req.Kode,
		"perusahaan_id": req.PerusahaanID,
	}
	CreateResponse(c, http.StatusOK, "success", "Barang created successfully.", barang)
}

func CreatePerusahaan(c *gin.Context) {
	var req models.Perusahaan
	var err error

	err = c.BindJSON(&req)
	if err != nil {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body.", nil)
		return
	}

	patternKode := "^[A-Z]{3}$"
	patternNoTelp := "^0[0-9]{9,12}$"

	var match bool

	match, err = regexp.MatchString(patternKode, req.Kode)
	if err != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to validate the kode.", nil)
		return
	}
	if !match {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid kode format. It should be 3 uppercase alphabet characters.", nil)
		return
	}

	match, err = regexp.MatchString(patternNoTelp, req.NoTelp)
	if err != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to validate no telp.", nil)
		return
	}
	if !match {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid no telp. It should have '0' as the prefix and be 10-13 digits in length.", nil)
		return
	}

	res := initializers.DB.Create(&req)
	if res.Error != nil {
		if strings.Contains(res.Error.Error(), "duplicate key") {
			CreateResponse(c, http.StatusConflict, "error", "Perusahaan with the certain ID is already exists.", nil)
		} else {
			CreateResponse(c, http.StatusInternalServerError, "error", "Failed to create perusahaan.", nil)
		}
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Perusahaan created successfully.", req)
}
