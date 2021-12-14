package controller

import (
	"resource-service/src/service"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	logger "resource-service/utils/logging"
	"resource-service/utils/wrapper/request"
	"resource-service/utils/wrapper/response"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

//ResourceController - Resource Controller
type ResourceController struct {
}

var log *zerolog.Logger = logger.GetInstance()

var resourceService = service.ResourceService{}

//AddResource - Controller to Add Resource
func (resource *ResourceController) AddResource(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestAddResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}

	acc_token := temp.tk

	//Calling AddResource Service
	msg, err := resourceService.AddResource(data, acc_token)
	if err != nil {

		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_ADD_RESOURCE, err)
		return // exit
	}
	response.SuccessWithMessage(c, msg)
}

//AddImageFile - Controlller to Add Image File
func (resource *ResourceController) AddImage(c *gin.Context) {

	//Fetch file from FormFile
	mpf, mpfh, err := c.Request.FormFile(constants.IMAGE)
	if err != nil {
		log.Error().Msg(constants.REQUIRED_REQUEST_BODY + ":" + err.Error())
		//Returns Error Code 500
		response.Send400(c, constants.REQUIRED_REQUEST_BODY)
		return // exit
	}
	defer mpf.Close()
	//Validate the Image Extension
	if request.ValidateImageExtension(c, mpfh.Filename) {
		return // exit
	}
	//Calling AddFile Service
	path, err := resourceService.AddImageAndFile(c, mpf, mpfh, constants.IMAGE_FOLDER)
	if err != nil {
		if err.Error() == constants.ERROR_FILE_NAME_IS_NOT_UNIQUE {
			log.Error().Msgf(constants.ERROR_FILE_NAME_IS_NOT_UNIQUE)
			//Return Error File Name is not Unique
			response.Send400(c, constants.ERROR_FILE_NAME_IS_NOT_UNIQUE)
			return // exit
		}
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_ADD_FILE, err)
		return // exit
	}
	response.SuccessWithData(c, path)
}

//AddFile - Controlller to Add PDF File
func (resource *ResourceController) AddFile(c *gin.Context) {

	//Fetch file from FormFile
	mpf, mpfh, err := c.Request.FormFile(constants.FILE)
	if err != nil {
		log.Error().Msg(constants.REQUIRED_REQUEST_BODY + ":" + err.Error())
		//Returns Error Code 500
		response.Send400(c, constants.REQUIRED_REQUEST_BODY)
		return // exit
	}
	defer mpf.Close()
	//Validate the File Extension
	if request.ValidateFileExtension(c, mpfh.Filename) {
		return // exit
	}
	//Calling AddFile Service
	path, err := resourceService.AddImageAndFile(c, mpf, mpfh, constants.FILE_FOLDER)
	if err != nil {
		if err.Error() == constants.ERROR_FILE_NAME_IS_NOT_UNIQUE {
			log.Error().Msgf(constants.ERROR_FILE_NAME_IS_NOT_UNIQUE)
			//Return Error File Name is not Unique
			response.Send400(c, constants.ERROR_FILE_NAME_IS_NOT_UNIQUE)
			return // exit
		}
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_ADD_FILE, err)
		return // exit
	}
	response.SuccessWithData(c, path)
}

//GetFile - Controller to Get File from bucket
func (resource *ResourceController) GetFile(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestGetFileDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}
	//Calling AddResource Service
	err := resourceService.GetFile(c, data)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_GET_FILE, err)
		return // exit
	}
	response.SuccessWithData(c, "")
}

//GetListOfResources - Controller to Get List of Resources
func (resource *ResourceController) GetListOfResources(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestGetListOfResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}
	//Calling GetListOfResources Service
	resources, err := resourceService.GetListOfResources(data)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_GET_LIST_OF_RESOURCES, err)
		return // exit
	}
	response.SuccessWithData(c, resources)
}

//GetResource - Controller to Get Resource
func (resource *ResourceController) GetResource(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestGetResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}

	acc_token := temp.tk
	//Calling GetListOfResources Service
	resources, err := resourceService.GetResource(data, acc_token)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_GET_RESOURCE, err)
		return // exit
	}
	response.SuccessWithData(c, resources)
}

//EditResource - Controller to Edit Resource
func (resource *ResourceController) EditResource(c *gin.Context) {
	//Bind json according to given structure
	var data dto.RequestEditResourceDTO
	if request.CheckJSON(c, &data) {
		return // exit
	}
	//Validator to Check Request Body is correct or not
	if request.CheckRequestBodyValidator(c, data) {
		return // exit
	}
	//Calling EditResource Service
	resources, err := resourceService.EditResource(data)
	if err != nil {
		//Returns Error Code 500
		response.Send500(c, constants.ERROR_IN_EDIT_RESOURCE, err)
		return // exit
	}

	response.SuccessWithMessage(c, resources)
}
