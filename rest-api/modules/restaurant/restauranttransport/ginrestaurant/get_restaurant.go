package ginrestaurant

import (
	"net/http"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantstore"
	"strconv"

	"github.com/gin-gonic/gin"

	"rest-api/component"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		getRestaurantBusiness := restaurantbusiness.NewGetRestaurantBusiness(store)

		data, err := getRestaurantBusiness.GetRestaurant(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
