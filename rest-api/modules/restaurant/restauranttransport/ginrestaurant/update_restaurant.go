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

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
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

		if err := updateRestaurantBusiness.UpdateRestaurant(c.Request.Context(), int(uid.GetLocationID()), &data); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
