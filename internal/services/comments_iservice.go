package service

import (
	"blog-system/internal/model"
	"blog-system/pkg/exception"
	"context"
)

type CommentService interface {
	// CRUD operations for Comment
	Create(ctx context.Context, req *model.CreateCommentReq) (*model.CreateCommentRes, *exception.Exception)
	Find(ctx context.Context, req *model.GetAllCommentsReq) (*model.GetAllCommentsRes, *exception.Exception)
	Detail(ctx context.Context, req *model.GetCommentByIDReq) (*model.GetCommentByIDRes, *exception.Exception)
	Update(ctx context.Context, req *model.UpdateCommentReq) (*model.UpdateCommentRes, *exception.Exception)
	Delete(ctx context.Context, req *model.DeleteCommentReq) (*model.DeleteCommentRes, *exception.Exception)
}
