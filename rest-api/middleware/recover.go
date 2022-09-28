package middleware

import (
	"github.com/gin-gonic/gin"
	"rest-api/common"
	"rest-api/component"
)

func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appError := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appError.StatusCode, appError)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
