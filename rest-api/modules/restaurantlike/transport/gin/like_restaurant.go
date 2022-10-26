package ginrestaurantlike

import (
	"net/http"
	"rest-api/modules/restaurant/restaurantstore"

	"github.com/gin-gonic/gin"

	"rest-api/common"
	"rest-api/component"
	restaurantlikebussiness "rest-api/modules/restaurantlike/business"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
	restaurantlikestorage "rest-api/modules/restaurantlike/store"
)

// POST /v1/restaurants/:id/like

func LikeRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet("user").(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantID: int(uid.GetLocationID()),
			UserID:       requester.GetUserID(),
		}

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		incrStore := restaurantstore.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantlikebussiness.NewLikeRestaurantBusiness(store, incrStore)

		if err := business.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
