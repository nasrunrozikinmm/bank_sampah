package route

import (
	"github.com/gin-gonic/gin"
	"github.com/iyunrozikin26/bank_sampah.git/src/config"
)

var router = gin.Default()

// func Run will start the server
func Run() {
	getRoutes()
	router.Run(":5000")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy

func getRoutes() {
	db := config.DB()
	v1 := router.Group("/api/v1")
	
	// daftarkan route lainnya
	addUserRoutes(v1, db)
}