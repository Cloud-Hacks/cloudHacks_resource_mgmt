package repository

import (
	"fmt"
	"reflect"
	"time"

	"resource-service/src/model"
	"resource-service/src/service/dto"
	"resource-service/utils/constants"
	"resource-service/utils/filter"
)

var res = []model.Resource{
	{ID: 1, Title: "SPIFEE", Category: "Tech", Status: "Published", Types: "PDF", Content: "FGY", FileLink: "/hyt.pdf", CreatedBy: 32, UpdatedBy: 54},
	{ID: 2, Title: "Devtron", Category: "CI/CD", Status: "Draft", Types: "TXT", Content: "DRE", FileLink: "/grew.txt", CreatedBy: 21, UpdatedBy: 12},
}

var img_resource = []model.Image{
	{ID: 1, ResourceId: 1, Link: "/gty.jpg", Status: "Draft", CreatedBy: 43, UpdatedBy: 23},
}

//ResourceRepository- Resource Repository
type ResourceRepository struct {
}

//AddResource - Store Resource into DB
func (resourceRepository *ResourceRepository) AddResource(data dto.RequestAddResourceDTO) (string, error) {
	//To store the resource
	var resource model.Resource
	resource.ID = data.ID
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
	// if err := database.GetInstance().Create(&resource).Error; err != nil {
	// 	return "", err
	// }

	res = append(res, resource)

	// c.IndentedJSON(http.StatusCreated, resource)

	var img_res = model.Image{}

	//Check Images Linked exist
	for _, link := range data.ImageLinks {
		//Create the record in Image Table
		img_res.ResourceId = resource.ID
		img_res.Link = link
		img_res.Status = constants.USED
		img_res.CreatedBy = data.CreatedBy
		img_res.CreatedAt = time.Now().UTC().Unix()
		img_res.UpdatedBy = data.CreatedBy
		img_res.UpdatedAt = time.Now().UTC().Unix()
	}
	img_resource = append(img_resource, img_res)

	return "Success", nil
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
	// if err := database.GetInstance().Select("id,title,category,status").Where("status != ? and title like \"%"+data.Filter.Query+"%\""+query, constants.DELETE).Limit(data.PageSize).Offset(data.PageNumber).Find(&resources).Error; err != nil {
	// 	return dto.ResponseGetListOfResourcesDTO{}, err
	// }
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
	// if err := database.GetInstance().Model(&model.Resource{}).Select("id").Count(&totalRecord).Error; err != nil {
	// 	return dto.ResponseGetListOfResourcesDTO{}, err
	// }
	//Return response
	return dto.ResponseGetListOfResourcesDTO{
		TotalRecord:     totalRecord,
		ListOfResources: listOfResources,
	}, nil
}

//GetResource - Get Resource based on ID from DB
func (resourceRepository *ResourceRepository) GetResource(data dto.RequestGetResourceDTO) (dto.ResponseGetResourceDTO, error) {

	for _, q := range res {
		for _, r := range img_resource {
			if q.ID == data.ID && q.ID == r.ResourceId {

				//Return response
				return dto.ResponseGetResourceDTO{
					ID:         data.ID,
					Title:      q.Title,
					Category:   q.Category,
					Status:     q.Status,
					Types:      q.Types,
					FileLink:   q.FileLink,
					Content:    q.Content,
					ImageLinks: r.Link,
					UpdatedBy:  q.UpdatedBy,
					UpdatedAt:  q.UpdatedAt,
				}, nil
			}
		}
	}
	return dto.ResponseGetResourceDTO{}, fmt.Errorf("No Data")
}

//EditResource - Edit Resource based on ID from DB
func (resourceRepository *ResourceRepository) EditResource(data dto.RequestEditResourceDTO) (string, error) {

	var image model.Image
	image.ResourceId = data.ID
	image.Link = data.ImageLinks
	image.Status = constants.USED
	image.CreatedBy = data.UpdatedBy
	image.CreatedAt = time.Now().UTC().Unix()
	image.UpdatedBy = data.UpdatedBy
	image.UpdatedAt = time.Now().UTC().Unix()

	return "Success", nil
}
