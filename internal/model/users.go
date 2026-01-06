package model

import (
	"blog-system/internal/entity"

	"github.com/google/uuid"
)

type BaseUsersReq struct {
	Username string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required,password,gte=8" example:"SecurePass123!"`
}

type CreateUserReq struct {
	BaseUsersReq
}

func (req BaseUsersReq) ToEntity(password string) *entity.Users {
	return &entity.Users{
		ReferencesId: uuid.NewString(),
		Username:     req.Username,
		Password:     password,
	}
}

type CreateUserRes struct {
	entity.Users
}

type LoginUserRes struct {
	Username string `json:"username" example:"john_doe"`
	Token    string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"` // JWT token example

}

type UpdateUserReq struct {
	BaseUsersReq
	ReferencesId string
}
type UpdateUserRes struct {
	entity.Users
}

type DeleteUserReq struct {
	BaseUsersReq
	ReferencesId string //uuid, will be get from query handler
}
type DeleteUserRes struct {
	ReferencesId string //uuid, will be get from query handler
}

type GetAllUsersReq struct {
	Page   PaginationParam
	Filter FilterParams
	Sort   OrderParam
}
type GetAllUsersRes struct {
	PaginationData[entity.Users]
}

type GetUserByIDReq struct {
	ReferencesId string `json:"references_id" name:"references_id"`
}

type GetUserByIDRes struct {
	entity.Users
}
