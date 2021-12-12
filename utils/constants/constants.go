package constants

//URL Path constants
const (
	LOGIN_USER_PATH            = "login-user"            //LOGIN_USER_PATH - URL Path to Login User
	ADD_RESOURCE_PATH          = "add-resource"          //AddResoucePath - URL Path to Add Resource
	ADD_IMAGE_PATH             = "add-image"             //ADD_IMAGE_PATH - URL Path to Add Image File into the bucket
	ADD_FILE_PATH              = "add-file"              //ADD_FILE_PATH - URL Path to Add File into the bucket
	GET_FILE_PATH              = "get-file"              //GET_FILE_PATH - URL Path to Get File from the bucket
	GET_LIST_OF_RESOURCES_PATH = "get-list-of-resources" //GET_LIST_OF_RESOURCES_PATH - URL Path to Get List of Resource
	GET_RESOURCE_PATH          = "view-resource"         //GET_RESOURCE_PATH - URL Path to Get Resource
	CHANGE_STATUS_PATH         = "change-status"         //CHANGE_STATUS_PATH - URL Path to Change Status of the resource
	EDIT_RESOURCE_PATH         = "edit-resource"         //EDIT_RESOURCE_PATH - URL Path to Edit Resource
	DELETE_RESOURCE_PATH       = "delete-resource"       //DELETE_RESOURCE_PATH - URL Path to Delete Resource
)

//Message constants
const (
	ERR_400                        = "err-400"
	ERR_404                        = "err-404"
	ERR_500                        = "err-500"
	INVALID_REQUEST_BODY           = "Invalid Request Body"
	REQUIRED_REQUEST_BODY          = "Required Request Body Missing"
	WRONG_REQUEST_BODY             = "Wrong Request Body : "
	SAME_STATUS_FOUND              = "Status is same as previous status"
	ERROR_IN_LOGIN_USER            = "Error in Login User"
	ERROR_IN_ADD_RESOURCE          = "Error in add Resource "
	ERROR_IN_ADD_IMAGE             = "Error in add Image "
	ERROR_IN_ADD_FILE              = "Error in add File "
	INVALID_EXTENSION              = "Invalid Extension : "
	ERROR_IN_GET_FILE              = "Error in Get File "
	ERROR_FILE_NAME_IS_NOT_UNIQUE  = "Error File Name is not Unique "
	ERROR_IN_GET_LIST_OF_RESOURCES = "Error in Get List of Resources "
	ERROR_IN_GET_RESOURCE          = "Error in Get Resource "
	ERROR_IN_STATUS_CHANGE         = "Error in Status Change "
	ERROR_IN_EDIT_RESOURCE         = "Error in Edit Resource "
	ERROR_IN_DELETE_RESOURCE       = "Error in Delete Resource "
)

//Common constants
const (
	USED           = "Used"
	FORWARD_SLASH  = "/"
	BACKWARD_SLASH = "\""
	DOT            = "."
	COMMA          = ","
	CLOSE_BRACKET  = ")"
	IMAGE_FOLDER   = "Image/"
	FILE_FOLDER    = "File/"
	IMAGE          = "image"
	FILE           = "file"
	CHARSET        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LENGTH         = 30
	DELETE         = "Delete"
	UN_USED        = "Un-Used"
)

const (
	CREDENTIALS_FILE = "conf/accuknox-website.json" //CREDENTIALS_FILE - Service Account Key in json Formate
)

const (
	AND_CATEGORY_IN = " and category in ("
	AND_STATUS      = " and status = "
)
