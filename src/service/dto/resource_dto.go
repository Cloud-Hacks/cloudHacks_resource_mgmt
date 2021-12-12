package dto

//RequestAddResourceDTO - DTO for add Resource Request
type RequestAddResourceDTO struct {
	ID         int      `json:"id" validate:"min=1"`
	Title      string   `json:"title" validate:"required"`
	Category   string   `json:"category" validate:"required"`
	Status     string   `json:"status" validate:"required,oneof= Draft Published"`
	Type       string   `json:"type" validate:"required,oneof= Text PDF"`
	ImageLinks []string `json:"image_links"`
	FileLink   string   `json:"file_link"`
	Content    string   `json:"content"`
	CreatedBy  int      `json:"created_by" validate:"required"`
}

//ResponseAddFileDTO - DTO for Add File Response
type ResponseAddFileDTO struct {
	FilePath string `json:"file_path"`
}

//RequestGetFileDTO - DTO for Get File Request
type RequestGetFileDTO struct {
	FilePath string `json:"file_path" validate:"required"`
}

//RequestGetListOfResourceDTO - DTO for Get List of Resource Request

type RequestGetListOfResourceDTO struct {
	PageNumber int       `json:"page_number" validate:"min=0"`
	PageSize   int       `json:"page_size" validate:"min=1"`
	Filter     FilterDTO `json:"filter"`
}

//FilterDTO - DTO for Filter Request
type FilterDTO struct {
	Category []string `json:"category"`
	Status   string   `json:"status"`
	Query    string   `json:"query"`
}

//ResponseGetListOfResourcesDTO - DTO for Get List of Resources Response
type ResponseGetListOfResourcesDTO struct {
	TotalRecord     int64               `json:"totalRecord"`
	ListOfResources []ListOfResourceDTO `json:"list_of_resources"`
}

//ListOfResourceDTO - DTO for List of Resource Response
type ListOfResourceDTO struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Status   string `json:"status"`
}

//RequestGetResourceDTO - DTO for Get Resource Request
type RequestGetResourceDTO struct {
	ID int `json:"id" validate:"min=0"`
}

type RequestChangeStatusDTO struct {
	ID        int    `json:"id" validate:"min=1"`
	Status    string `json:"status" validate:"required,oneof= Draft Published"`
	UpdatedBy int    `json:"updated_by" validate:"required"`
}

//ResponseGetListOfResourcesDTO - DTO for Get Resource Response
type ResponseGetResourceDTO struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Category   string `json:"category"`
	Status     string `json:"status"`
	Types      string `json:"types"`
	ImageLinks string `json:"image_links"`
	FileLink   string `json:"file_link"`
	Content    string `json:"content"`
	UpdatedBy  int    `json:"updated_by"`
	UpdatedAt  int64  `json:"updated_at"`
}

//RequestEditResourceDTO - DTO for Edit Resource
type RequestEditResourceDTO struct {
	ID         int    `json:"id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Category   string `json:"category" validate:"required"`
	Type       string `json:"type" validate:"oneof= Text PDF"`
	Status     string `json:"status" validate:"oneof= Draft Published"`
	Content    string `json:"content"`
	ImageLinks string `json:"image_links"`
	FileLink   string `json:"file_link"`
	UpdatedBy  int    `json:"updated_by" validate:"required"`
}

//RequestDeleteResourceDTO - DTO for Delete Resource
type RequestDeleteResourceDTO struct {
	ID        int `json:"id" validate:"required"`
	UpdatedBy int `json:"updated_by" validate:"required"`
}
