package main

import (
	"fmt"
	"labpro/single-service/controllers"
	"labpro/single-service/initializers"
	// "labpro/single-service/internal/browser"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadConfig()
	initializers.ConnectDB(true)
	fmt.Println("Config initialized and DB connected.")
}

// @title Single-Service API documentation
// @version 1.0.0
// @host localhost:3002
// @BasePath /
func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8000", "http://localhost:5173", "https://ohl-fe.vercel.app"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Accept", "Accept-Encoding", "Content-Length", "X-CSRF-Token", "Authorization"},
	}))
	
	// POST
	router.POST("/login", controllers.Login)
	router.POST("/barang", controllers.Auth, controllers.CreateBarang) //checked
	router.POST("/perusahaan", controllers.Auth, controllers.CreatePerusahaan) //checked
	
	// GET
	router.GET("/self", controllers.GetSelf)
	router.GET("/barang", controllers.GetBarang)
	router.GET("/barang/:id", controllers.GetBarangById)
	router.GET("/perusahaan", controllers.GetPerusahaan)
	router.GET("/perusahaan/:id", controllers.GetPerusahaanById) //checked

	// DELETE
	router.DELETE("/barang/:id", controllers.Auth, controllers.DeleteBarangById)
	router.DELETE("/perusahaan/:id", controllers.Auth, controllers.DeletePerusahaanById)
	
	// UPDATE
	router.PUT("/barang/:id", controllers.Auth, controllers.UpdateBarangById)
	router.PUT("/perusahaan/:id", controllers.Auth, controllers.UpdatePerusahaanById)
	
	URL := "0.0.0.0:" + initializers.Cfg.PORT
	// status := browser.Open("http://" + URL)
	// if !status {
	// 	fmt.Println("Failed to open browser.")
	// }
	router.Run(URL)
}