package response

import (
	"net/http"

	"resource-service/src/service/dto"

	"github.com/gin-gonic/gin"
)

//SuccessWithMessage - Return with Success Message
func SuccessWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, dto.SuccessDTO{
		Message: msg,
	})
}

//SuccessWithData - Return expected Data
func SuccessWithData(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{"jwt": msg})

}
