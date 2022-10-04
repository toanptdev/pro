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
			panic(err)
		}

		requester := c.MustGet("user").(common.Requester)

		data.OwnerID = requester.GetUserID()

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurantBusiness(store)
		if err := business.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.GenerateUID(common.DBTypeRestaurant)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
