package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"rest-api/common"
	"rest-api/component"
)

func GetProfile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet("user").(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
