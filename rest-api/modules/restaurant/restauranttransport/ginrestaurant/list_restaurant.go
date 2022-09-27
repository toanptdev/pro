package ginrestaurant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"

	"rest-api/modules/restaurant/restaurantbusiness"
	"rest-api/modules/restaurant/restaurantmodel"
	"rest-api/modules/restaurant/restaurantstore"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		paging.Fulfill()

		store := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantbusiness.NewListRestaurantBusiness(store)
		result, err := business.ListRestaurant(c.Request.Context(), &filter, &paging)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResp(result, paging, filter))
	}
}
