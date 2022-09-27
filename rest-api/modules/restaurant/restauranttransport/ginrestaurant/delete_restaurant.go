package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantstore"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		deleteRestaurantBusiness := restaurantbusiness.NewDeleteRestaurantStore(store)

		if err := deleteRestaurantBusiness.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
