//login_controller

package controller

import (
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	"resource-service/utils/wrapper/request"
	"resource-service/utils/wrapper/response"

	"github.com/gin-gonic/gin"
)

//LoginController - Login Controller
type LoginController struct {
	tk string
}

var temp LoginController

// var userService = &service.UserService{}

//LoginUser - Constroller to Login User

func (user *LoginController) LoginUser(c *gin.Context) {

	//Bind json according to given structure
	var data dto.RequestUserDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}

	dt, err := request.CheckAuthorisedUser(c, data)

	temp.tk = dt

	// Returns Error Code 404
	if err != nil {
		response.Send404(c, constants.ERROR_IN_LOGIN_USER, err)
		return
	}

	response.SuccessWithData(c, dt)
}
