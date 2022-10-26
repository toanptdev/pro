package ginrestaurantlike

import (
	"net/http"
	"rest-api/modules/restaurant/restaurantstore"

	"github.com/gin-gonic/gin"

	"rest-api/common"
	"rest-api/component"
	restaurantlikebussiness "rest-api/modules/restaurantlike/business"
	restaurantlikestorage "rest-api/modules/restaurantlike/store"
)

// POST /v1/restaurants/:id/unlike

func UnLikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet("user").(common.Requester)

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		decrStore := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantlikebussiness.NewLUnLikeRestaurantBusiness(store, decrStore)

		if err := business.UnLikeRestaurant(c.Request.Context(), requester.GetUserID(), int(uid.GetLocationID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
