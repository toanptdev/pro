package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"

	"rest-api/component"
	hasher "rest-api/component/hasher"
	"rest-api/component/tokenprovider/jwt"
	userbusiness "rest-api/modules/users/business"
	usermodel "rest-api/modules/users/model"
	userstorage "rest-api/modules/users/store"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data usermodel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSqlStore(appCtx.GetMainDBConnection())
		tokenProvider := jwt.NewJWTProvider(appCtx.GetSecret())
		hash := hasher.NewMd5Hash()
		tokenConfig, err := appCtx.NewTokenConfig()
		if err != nil {
			panic(err)
		}
		loginBusiness := userbusiness.NewLoginBusiness(store, tokenProvider, hash, tokenConfig)
		account, err := loginBusiness.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
