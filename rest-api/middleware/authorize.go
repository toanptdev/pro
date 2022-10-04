package middleware

import (
	"errors"
	"rest-api/component/tokenprovider/jwt"

	"github.com/gin-gonic/gin"
	"rest-api/common"
	"rest-api/component"
	userstorage "rest-api/modules/users/store"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader
	}

	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewJWTProvider(appCtx.GetSecret())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSqlStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserID})
		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}
		user.Mask()
		c.Set("user", user)
		c.Next()
	}
}

var ErrWrongAuthHeader = common.NewCustomError(
	errors.New("wrong authen header"),
	"wrong authen header",
	"ErrWrongAuthHeader",
)
