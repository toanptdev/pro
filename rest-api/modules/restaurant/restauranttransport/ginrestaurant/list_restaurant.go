package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/modules/restaurant/restaurantrepository"
	restaurantlikestorage "rest-api/modules/restaurantlike/store"

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
		likeStore := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		repository := restaurantrepository.NewRestaurantLikeRepository(store, likeStore)
		business := restaurantbusiness.NewListRestaurantBusiness(repository)
		result, err := business.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		for i := range result {
			result[i].Mask(false)
		}

	}
}
