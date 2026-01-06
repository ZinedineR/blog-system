package service

import (
	"blog-system/internal/model"
	"blog-system/pkg/exception"
	"context"
)

type PostService interface {
	// CRUD operations for Post
	Create(ctx context.Context, req *model.CreatePostReq) (*model.CreatePostRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllPostsReq) (*model.GetAllPostsRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetPostByIDReq) (*model.GetPostByIDRes, *exception.Exception)
	Update(ctx context.Context, req *model.UpdatePostReq) (*model.UpdatePostRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeletePostReq) (*model.DeletePostRes, *exception.Exception)
}
