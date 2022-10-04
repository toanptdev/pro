package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
	"rest-api/component/hasher"
	userbusiness "rest-api/modules/users/business"
	usermodel "rest-api/modules/users/model"
	userstorage "rest-api/modules/users/store"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		userBusiness := userbusiness.NewRegisterBusiness(store, md5)

		if err := userBusiness.Register(c.Request.Context(), data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
