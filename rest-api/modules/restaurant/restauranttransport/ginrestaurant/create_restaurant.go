package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"

	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)
		if err := business.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
