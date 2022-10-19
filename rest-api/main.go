package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-api/component"
	"rest-api/middleware"
	ginrestaurantlike "rest-api/modules/restaurantlike/transport/gin"
	"rest-api/modules/users/transport/ginuser"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"rest-api/modules/restaurant/restauranttransport/ginrestaurant"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	fmt.Println(dsn)
	secret := os.Getenv("secret")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	runService(db, secret)
}

func runService(db *gorm.DB, secret string) error {
	appContext := component.NewAppContext(db, secret)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/login", ginuser.Login(appContext))
	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.GetProfile(appContext))

	restaurant := v1.Group("/restaurants", middleware.RequireAuth(appContext))
	{
		restaurant.POST("", ginrestaurant.CreateRestaurant(appContext))
		restaurant.GET("", ginrestaurant.ListRestaurant(appContext))
		restaurant.GET(":id", ginrestaurant.GetRestaurant(appContext))
		restaurant.PATCH(":id", ginrestaurant.UpdateRestaurant(appContext))
		restaurant.DELETE(":id", ginrestaurant.DeleteRestaurant(appContext))

		restaurant.GET("/:id//liked-users", ginrestaurantlike.ListUsers(appContext))
	}

	return r.Run()
}
