package dto

//UserDTO - DTO for add Resource Request
type RequestUserDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//ResponseLogin dto
type ResponseLogin struct {
	JWT string `json:"access_token"`
}
