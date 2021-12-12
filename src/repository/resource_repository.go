package repository

import (
	"errors"
	"reflect"
	"time"

	"resource-service/src/model"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	"resource-service/utils/database"
	"resource-service/utils/filter"
)

//ResourceRepository- Resource Repository
type ResourceRepository struct {
}

//AddResource - Store Resource into DB
func (resourceRepository *ResourceRepository) AddResource(data dto.RequestAddResourceDTO) (string, error) {
	//To store the resource
	var resource model.Resource
	resource.Title = data.Title
	resource.Category = data.Category
	resource.Status = data.Status
	resource.Types = data.Type
	resource.Content = data.Content
	resource.FileLink = data.FileLink
	resource.CreatedBy = data.CreatedBy
	resource.CreatedAt = time.Now().UTC().Unix()
	resource.UpdatedBy = data.CreatedBy
	resource.UpdatedAt = time.Now().UTC().Unix()
	//Create the record in Resource Table
	if err := database.GetInstance().Create(&resource).Error; err != nil {
		return "", err
	}
	//Check Images Linked exist
	for _, link := range data.ImageLinks {
		//Create the record in Image Table
		if err := database.GetInstance().Create(&model.Image{
			ResourceId: resource.ID,
			Link:       link,
			Status:     constants.USED,
			CreatedBy:  data.CreatedBy,
			CreatedAt:  time.Now().UTC().Unix(),
			UpdatedBy:  data.CreatedBy,
			UpdatedAt:  time.Now().UTC().Unix(),
		}).Error; err != nil {
			return "", err
		}

	}
	return "Success", nil
}

//UniqureFileName - Check File Name and Path is Uniqure or not
func (resourceRepository *ResourceRepository) UniqureFileName(folder, filePath string) bool {

	switch folder {
	case constants.IMAGE_FOLDER:
		if err := database.GetInstance().Select("id").Where("link = ? ", folder+filePath).First(&model.Image{}).Error; err != nil {
			return false
		}
	case constants.FILE_FOLDER:
		if err := database.GetInstance().Select("id").Where("file_link = ? ", folder+filePath).First(&model.Resource{}).Error; err != nil {
			return false
		}
	}
	return true
}

//GetListOfResources - Get List of Resources from DB
func (resourceRepository *ResourceRepository) GetListOfResources(data dto.RequestGetListOfResourceDTO) (dto.ResponseGetListOfResourcesDTO, error) {

	//Store the List of Resources as Output
	var listOfResources []dto.ListOfResourceDTO
	//Store total Record in Resource
	var totalRecord int64
	//Store Filter queries
	var query string
	// Check Filter is exist or not
	if !reflect.DeepEqual(data.Filter, (dto.FilterDTO{})) {
		//Check Category Filter Exist or not
		if len(data.Filter.Category) != 0 {
			//Fetch the query
			query = filter.QueryFilterString(data.Filter.Category, constants.AND_CATEGORY_IN)
		}
		//Check Status Filter Exist or not
		if data.Filter.Status != "" {
			//Add query to filter by status
			query += constants.AND_STATUS + "\"" + data.Filter.Status + "\""
		}
	}

	//Store the resources
	var resources []model.Resource
	//Fetch the resources from Resource Table
	if err := database.GetInstance().Select("id,title,category,status").Where("status != ? and title like \"%"+data.Filter.Query+"%\""+query, constants.DELETE).Limit(data.PageSize).Offset(data.PageNumber).Find(&resources).Error; err != nil {
		return dto.ResponseGetListOfResourcesDTO{}, err
	}
	//Looping resources
	for _, resource := range resources {
		//Append data into list of Resources
		listOfResources = append(listOfResources, dto.ListOfResourceDTO{
			ID:       resource.ID,
			Title:    resource.Title,
			Category: resource.Category,
			Status:   resource.Status,
		})
	}
	//Fetch the Total Record from Resource Table
	if err := database.GetInstance().Model(&model.Resource{}).Select("id").Count(&totalRecord).Error; err != nil {
		return dto.ResponseGetListOfResourcesDTO{}, err
	}
	//Return response
	return dto.ResponseGetListOfResourcesDTO{
		TotalRecord:     totalRecord,
		ListOfResources: listOfResources,
	}, nil
}

//GetResource - Get Resource based on ID from DB
func (resourceRepository *ResourceRepository) GetResource(data dto.RequestGetResourceDTO) (dto.ResponseGetResourceDTO, error) {
	//Store resource
	var resource model.Resource
	//Fetch the resources from Resource Table
	if err := database.GetInstance().Select("title,category,status,types,content,file_link,updated_by,updated_at").Where("id = ? and status != ?", data.ID, constants.DELETE).First(&resource).Error; err != nil {
		return dto.ResponseGetResourceDTO{}, err
	}
	//Store Image Links
	var imageLinks []string
	//Fetch the Image Links associated Resource ID
	if err := database.GetInstance().Model(&model.Image{}).Select("link").Where("resource_id = ? and status != ?", data.ID, constants.UN_USED).Scan(&imageLinks).Error; err != nil {
		return dto.ResponseGetResourceDTO{}, err
	}

	//Return response
	return dto.ResponseGetResourceDTO{
		ID:         data.ID,
		Title:      resource.Title,
		Category:   resource.Category,
		Status:     resource.Status,
		Types:      resource.Types,
		FileLink:   resource.FileLink,
		Content:    resource.Content,
		ImageLinks: imageLinks,
		UpdatedBy:  resource.UpdatedBy,
		UpdatedAt:  resource.UpdatedAt,
	}, nil
}

//ChangeStatus - Change Resource Status based on ID from DB
func (resourceRepository *ResourceRepository) ChangeStatus(data dto.RequestChangeStatusDTO) (string, error) {

	//Store resource
	var resource model.Resource
	//Fetch the resources from Resource Table
	if err := database.GetInstance().Select("status").Where("id = ? and status != ?", data.ID, constants.DELETE).First(&resource).Error; err != nil {
		return "", err
	}

	//Change the status of the resources from Resource Table if status is not same
	if resource.Status != data.Status {
		if err := database.GetInstance().Where("id = ?", data.ID).UpdateColumns(&model.Resource{
			Status: data.Status, UpdatedBy: data.UpdatedBy}).Error; err != nil {
			return "", err
		}
	} else {
		return "", errors.New(constants.SAME_STATUS_FOUND)
	}

	//Return response
	return "Success", nil
}

//EditResource - Edit Resource based on ID from DB
func (resourceRepository *ResourceRepository) EditResource(data dto.RequestEditResourceDTO) (string, error) {
	//Check Resource exist or not
	if err := database.GetInstance().Select("id").Where("id = ? and status != ?", data.ID, constants.DELETE).First(&model.Resource{}).Error; err != nil {
		return "", err
	}
	//Edit Resource
	if err := database.GetInstance().Where("id = ? and status != ?", data.ID, constants.DELETE).UpdateColumns(&model.Resource{
		Title:     data.Title,
		Category:  data.Category,
		Types:     data.Type,
		Status:    data.Status,
		Content:   data.Content,
		FileLink:  data.FileLink,
		UpdatedBy: data.UpdatedBy,
	}).Error; err != nil {
		return "", err
	}
	//Check Images Linked exist
	for _, link := range data.ImageLinks {
		//To store imageLinks
		var image model.Image
		image.ResourceId = data.ID
		image.Link = link
		image.Status = constants.USED
		image.CreatedBy = data.UpdatedBy
		image.CreatedAt = time.Now().UTC().Unix()
		image.UpdatedBy = data.UpdatedBy
		image.UpdatedAt = time.Now().UTC().Unix()
		//Update if Record already exist otherwise create New Record
		if err := database.GetInstance().Where("resource_id = ? and status != ?", data.ID, constants.DELETE).Assign(model.Image{Link: link}).FirstOrCreate(&image).Error; err != nil {
			return "", nil
		}
	}
	return "Success", nil
}

//DeleteResource - Delete Resource based on ID from DB
func (resourceRepository *ResourceRepository) DeleteResource(data dto.RequestDeleteResourceDTO) (string, error) {

	//Check Resource exist or not
	if err := database.GetInstance().Select("id").Where("id = ? and status != ?", data.ID, constants.DELETE).First(&model.Resource{}).Error; err != nil {
		return "", err
	}
	//Delete Resource
	if err := database.GetInstance().Where("id = ? ", data.ID).UpdateColumns(&model.Resource{
		Status:    constants.DELETE,
		UpdatedBy: data.UpdatedBy,
	}).Error; err != nil {
		//return
		return "", err
	}
	//Delete linked images
	if err := database.GetInstance().Where("resource_id = ? ", data.ID).UpdateColumns(&model.Image{
		Status:    constants.UN_USED,
		UpdatedBy: data.UpdatedBy,
	}).Error; err != nil {
		//return
		return "", err
	}
	//return
	return "Success", nil
}
