package controllers

import (
	"fmt"
	"labpro/single-service/initializers"
	"labpro/single-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBarangById(c *gin.Context) {
	id := c.Param("id")
	var barang models.Barang

	res := initializers.DB.First(&barang, "id = ?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to get barang.", nil)
		return
	}

	res = initializers.DB.Delete(&models.Barang{}, "id = ?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to delete barang.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Barang has been deleted.", barang)
}

func DeletePerusahaanById(c *gin.Context) {
	id := c.Param("id")
	var perusahaan models.Perusahaan

	res := initializers.DB.First(&perusahaan, "id = ?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to get perusahaan.", nil)
		return
	}

	res = initializers.DB.Delete(&models.Perusahaan{}, "id = ?", id)
	if res.Error != nil {
		fmt.Println(res.Error)
		CreateResponse(c, http.StatusInternalServerError, "error", "Failed to delete perusahaan.", nil)
		return
	}

	CreateResponse(c, http.StatusOK, "success", "Perusahaan has been deleted.", perusahaan)
}