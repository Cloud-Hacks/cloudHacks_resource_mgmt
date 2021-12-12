package request

import (
	"strings"

	"net/http"
	"time"

	"resource-service/src/model"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	logger "resource-service/utils/logging"
	"resource-service/utils/wrapper/response"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

var usrdt = []model.User{
	{ID: 1, FirstName: "Deter", LastName: "Fir", Email: "ety@k1.co", Password: "poetry@", Role: "Dev"},
	{ID: 2, FirstName: "Peter", LastName: "Ge", Email: "frt@k1.co", Password: "yutr@12", Role: "Art"},
}

var log *zerolog.Logger = logger.GetInstance()

var ACCESS_SECRET = viper.GetString("access_secret.code")

type Claims struct {
	UID int `json:"UserId"`
	jwt.StandardClaims
}

var claims = &Claims{}

var dt model.User

//Check JSON according to given structure
func CheckJSON(c *gin.Context, data interface{}) bool {
	if err := c.BindJSON(&data); err != nil {
		log.Error().Msgf(constants.WRONG_REQUEST_BODY + err.Error())
		//Return Error Invalid Request Body
		response.Send400(c, constants.INVALID_REQUEST_BODY)
		return true
	}
	return false
}

//Validate Data According to Request Body
func CheckRequestBodyValidator(c *gin.Context, data interface{}) bool {
	validation := validator.New()
	if err := validation.Struct(data); err != nil {
		log.Error().Msgf(constants.WRONG_REQUEST_BODY + err.Error())
		//Return Error Required Request Body
		response.Send400(c, constants.REQUIRED_REQUEST_BODY)
		return true
	}
	return false
}

//Validate the Image Extension
func ValidateImageExtension(c *gin.Context, filename string) bool {
	//Fetch the image extension
	extension := filename[strings.LastIndex(filename, ".")+1:]
	//Check Image extension is valid or not
	switch extension {
	case "jpg", "jpeg", "png", "svg":
	default:
		//Return when extension is invalid
		log.Error().Msgf(constants.INVALID_EXTENSION + extension)
		response.Send400(c, constants.INVALID_EXTENSION+extension)
		return true
	}
	return false
}

//Validate the File Extension
func ValidateFileExtension(c *gin.Context, filename string) bool {
	//Fetch the file extension
	extension := filename[strings.LastIndex(filename, ".")+1:]
	//Check File extension is valid or not
	switch extension {
	case "pdf":
	default:
		///Return when extension is invalid
		log.Error().Msgf(constants.INVALID_EXTENSION + extension)
		response.Send400(c, constants.INVALID_EXTENSION+extension)
		return true
	}
	return false
}

//Authenticate and Authorise the user request with the provided credentials
func CheckAuthorisedUser(c *gin.Context, data dto.RequestUserDTO) (interface{}, error) {

	// Encode the password
	// hash := sha512.New()
	// hash.Write([]byte(data.Email))
	// hash.Write([]byte(data.Password))
	// pwd := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	for _, u := range usrdt {
		if u.Email == data.Email && u.Password == data.Password {

			// if err := database.GetInstance().Where("email = ? AND password = ?", data.Email, data.Password).First(&dt).Error; err != nil {
			// 	return " ", err
			// }

			//update the last_login
			// dt.LastLogin = time.Now().UTC().Unix()
			// database.GetInstance().Save(&dt)

			expirationTime := time.Now().Add(90 * time.Second)

			//access the token with the claims used for signing
			token, err := AccessToken(dt.ID, dt.Email, expirationTime)

			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return " ", err
			}

			usrdt := GetResData(token)

			return usrdt, nil
		}
	}

	return "Null User Data", nil

}

// Access token from the user details
func AccessToken(userId int, userEmail string, expTime time.Time) (string, error) {

	claims := Claims{
		UID: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    userEmail,
			ExpiresAt: expTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := at.SignedString([]byte(ACCESS_SECRET))

	return token, err
}

// Get the User Data to response user
func GetResData(token string) interface{} {

	UsrDt := &dto.ResponseLogin{
		JWT: token,
	}

	return UsrDt
}
