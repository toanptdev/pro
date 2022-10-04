package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/common"
	"rest-api/component"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet("user").(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
