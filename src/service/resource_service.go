package service

import (
	"errors"
	"io"
	"mime/multipart"

	"resource-service/src/repository"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

//ResourceService - Resource Service
type ResourceService struct {
}

var resourceRepository repository.ResourceRepository = repository.ResourceRepository{}

//AddResource - Service to Add Resource
func (resource *ResourceService) AddResource(data dto.RequestAddResourceDTO) (string, error) {
	//Calling AddResource Repository
	msg, err := resourceRepository.AddResource(data)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (resource *ResourceService) AddImageAndFile(c *gin.Context, mpf multipart.File, mpfh *multipart.FileHeader, folder string) (dto.ResponseAddFileDTO, error) {

	//Check File Name is uniqure or not
	if resourceRepository.UniqureFileName(folder, mpfh.Filename) {
		return dto.ResponseAddFileDTO{}, errors.New(constants.ERROR_FILE_NAME_IS_NOT_UNIQUE)
	}

	//Define Bucket Name
	bucket := viper.GetString("bucket.name")
	//Define appengine context
	ctx := appengine.NewContext(c.Request)
	//Create Google Cloud Storage Client
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(constants.CREDENTIALS_FILE))
	if err != nil {
		return dto.ResponseAddFileDTO{}, err
	}
	//Close Google Cloud Storage Client
	defer client.Close()
	//Create storage Writer
	writer := client.Bucket(bucket).Object(folder + mpfh.Filename).NewWriter(ctx)
	//Copy file into location
	if _, err := io.Copy(writer, mpf); err != nil {
		return dto.ResponseAddFileDTO{}, err
	}
	//Close Writer
	if err := writer.Close(); err != nil {
		return dto.ResponseAddFileDTO{}, err
	}
	//Return File Path
	return dto.ResponseAddFileDTO{
		FilePath: writer.Name,
	}, nil
}

//GetFile - Service to Get File from Bucket
func (resource *ResourceService) GetFile(c *gin.Context, data dto.RequestGetFileDTO) error {
	//Define Bucket Name
	bucket := viper.GetString("bucket.name")
	//Define appengine context
	ctx := appengine.NewContext(c.Request)
	//Create Google Cloud Storage Client
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(constants.CREDENTIALS_FILE))
	if err != nil {
		return err
	}
	//Close Google Cloud Storage Client
	defer client.Close()
	//Create storage Reader
	reader, err := client.Bucket(bucket).Object(data.FilePath).NewReader(ctx)
	if err != nil {
		return err
	}
	//Close Reader
	defer reader.Close()
	//Copy Reader to context Writer
	if _, err := io.Copy(c.Writer, reader); err != nil {
		return err
	}
	return nil
}

//GetListOfResources - Service to Get List of Resources
func (resource *ResourceService) GetListOfResources(data dto.RequestGetListOfResourceDTO) (dto.ResponseGetListOfResourcesDTO, error) {

	//Calling GetListOfResources Repository
	resources, err := resourceRepository.GetListOfResources(data)
	if err != nil {
		return dto.ResponseGetListOfResourcesDTO{}, err
	}

	return resources, nil
}

//GetResource - Service to Get Resource
func (resource *ResourceService) GetResource(data dto.RequestGetResourceDTO) (dto.ResponseGetResourceDTO, error) {

	//Calling GetListOfResources Repository
	resources, err := resourceRepository.GetResource(data)
	if err != nil {
		return dto.ResponseGetResourceDTO{}, err
	}

	return resources, nil
}

//ChangeStatus - Service to Change Resource Status
func (resource *ResourceService) ChangeStatus(data dto.RequestChangeStatusDTO) (string, error) {

	//Calling ChangeStatus Repository
	msg, err := resourceRepository.ChangeStatus(data)

	if err != nil {
		return "", err
	}

	return msg, nil
}

//EditResource - Service to Edit Resource
func (resource *ResourceService) EditResource(data dto.RequestEditResourceDTO) (string, error) {

	//Calling EditResource Repository
	resources, err := resourceRepository.EditResource(data)
	if err != nil {
		return "", err
	}

	return resources, nil
}

//DeleteResource - Service to Delete Resource
func (resource *ResourceService) DeleteResource(data dto.RequestDeleteResourceDTO) (string, error) {

	//Calling DeleteResource Repository
	resources, err := resourceRepository.DeleteResource(data)
	if err != nil {
		return "", err
	}

	return resources, nil
}
