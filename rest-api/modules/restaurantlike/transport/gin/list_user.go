package ginrestaurantlike

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	restaurantlikebussiness "rest-api/modules/restaurantlike/business"
	restaurantlikemodel "rest-api/modules/restaurantlike/model"
	restaurantlikestorage "rest-api/modules/restaurantlike/store"
)

func ListUsers(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//var filter restaurantlikemodel.Filter
		//
		//if err := c.ShouldBind(&filter); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": err,
		//	})
		//	return
		//}

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		filter := restaurantlikemodel.Filter{
			RestaurantID: int(uid.GetLocationID()),
		}

		fmt.Println("filter: ", filter)

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		paging.Fulfill()

		store := restaurantlikestorage.NewSqlStore(appCtx.GetMainDBConnection())
		business := restaurantlikebussiness.NewListUserLikeRestaurantBusiness(store)
		result, err := business.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResp(result, paging, filter))
	}
}
