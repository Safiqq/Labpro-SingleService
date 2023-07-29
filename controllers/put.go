package controllers

import (
	"labpro/single-service/initializers"
	"labpro/single-service/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func UpdateBarangById(c *gin.Context) {
	id := c.Param("id")

	var req models.Barang
	if err := c.BindJSON(&req); err != nil {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body.", nil)
		return
	}

	if req.Harga == 0 {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body. Harga can't be 0.", nil)
		return
	}

	var barang models.Barang
	res := initializers.DB.First(&barang, "id = ?", id)
	if res.Error != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve barang.", nil)
		return
	}

	barang.Nama = req.Nama
	barang.Harga = req.Harga
	barang.Stok = req.Stok
	barang.Kode = req.Kode
	barang.PerusahaanID = req.PerusahaanID

	res = initializers.DB.Save(&barang)
	if res.Error != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to update barang.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Barang updated successfully.", barang)
}

func UpdatePerusahaanById(c *gin.Context) {
	id := c.Param("id")

	var req models.Perusahaan
	if err := c.BindJSON(&req); err != nil {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid request body.", nil)
		return
	}

	patternNoTelp := "^0[0-9]{9,12}$"
	
	match, err := regexp.MatchString(patternNoTelp, req.NoTelp)
	if err != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to validate no telp.", nil)
		return
	}
	if !match {
		CreateResponse(c, http.StatusBadRequest, "error", "Invalid no telp. It should have '0' as the prefix and be 10-13 digits in length.", nil)
		return
	}

	var perusahaan models.Perusahaan
	res := initializers.DB.First(&perusahaan, "id = ?", id)
	if res.Error != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to retrieve perusahaan.", nil)
		return
	}

	perusahaan.Nama = req.Nama
	perusahaan.Alamat = req.Alamat
	perusahaan.NoTelp = req.NoTelp
	perusahaan.Kode = req.Kode

	res = initializers.DB.Save(&perusahaan)
	if res.Error != nil {
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to update perusahaan.", nil)
		return
	}
	
	CreateResponse(c, http.StatusOK, "success", "Perusahaan updated successfully.", perusahaan)
}