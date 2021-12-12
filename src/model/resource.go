package model

//Resource - Model for Resources Table
type Resource struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	Types     string `json:"types"`
	Content   string `json:"content"`
	FileLink  string `json:"file_link"`
	CreatedBy int    `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
	UpdatedBy int    `json:"updated_by"`
	UpdatedAt int64  `json:"updated_at"`
}

//Image - Model for Images Table
type Image struct {
	ID         int    `json:"id"`
	ResourceId int    `json:"resource_id"`
	Link       string `json:"link"`
	Status     string `json:"status"`
	CreatedBy  int    `json:"created_by"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedBy  int    `json:"updated_by"`
	UpdatedAt  int64  `json:"updated_at"`
}
