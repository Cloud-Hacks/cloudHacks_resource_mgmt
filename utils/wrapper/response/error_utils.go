package response

import (
	"net/http"

	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	logger "resource-service/utils/logging"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var log *zerolog.Logger = logger.GetInstance()

// Return Error After service function called
func Send500(c *gin.Context, ErrMsg string, err error) {
	log.Error().Msg(ErrMsg + err.Error())
	c.JSON(http.StatusBadGateway, dto.ErrorDTO{
		ErrorCode:    constants.ERR_500,
		ErrorMessage: ErrMsg,
	})
}

// checkjsonbind and validation
func Send400(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, dto.ErrorDTO{
		ErrorCode:    constants.ERR_400,
		ErrorMessage: msg,
	})
}

// Return Error After Token Expires
func Send404(c *gin.Context, ErrMsg string, err error) {
	log.Error().Msgf(ErrMsg, err.Error())
	c.JSON(http.StatusNotFound, dto.ErrorDTO{
		ErrorCode:    constants.ERR_404,
		ErrorMessage: err.Error(),
	})
}
