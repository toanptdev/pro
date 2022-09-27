package main

import (
	"log"
	"net/http"
	"os"
	"rest-api/component"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"rest-api/modules/restaurant/restauranttransport/ginrestaurant"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	runService(db)
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appContext := component.NewAppContext(db)
	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))
		restaurant.GET("", ginrestaurant.ListRestaurant(appContext))
		restaurant.GET(":id", ginrestaurant.GetRestaurant(appContext))
		restaurant.PATCH(":id", ginrestaurant.UpdateRestaurant(appContext))
		restaurant.DELETE(":id", ginrestaurant.DeleteRestaurant(appContext))
	}

	return r.Run()
}