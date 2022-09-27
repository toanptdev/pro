package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
			return
		}

		var data restaurantmodel.RestaurantUpdates
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
			return
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		updateRestaurantBusiness := restaurantbusiness.NewUpdateRestaurantStore(store)

		if err := updateRestaurantBusiness.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
