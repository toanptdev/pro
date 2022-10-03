package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantstore"

	"rest-api/component"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		getRestaurantBusiness := restaurantbusiness.NewGetRestaurantBusiness(store)

		data, err := getRestaurantBusiness.GetRestaurant(c.Request.Context(), int(uid.GetLocationID()))
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
