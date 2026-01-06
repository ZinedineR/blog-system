package service

import (
	"blog-system/internal/model"
	"blog-system/pkg/exception"
	"context"
)

type UserService interface {
	// CRUD operations for User
	Register(
		ctx context.Context, req *model.CreateUserReq,
	) (*model.CreateUserRes, *exception.Exception)
	Login(ctx context.Context, req *model.CreateUserReq) (*model.LoginUserRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllUsersReq) (*model.GetAllUsersRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetUserByIDReq) (*model.GetUserByIDRes, *exception.Exception)
}
