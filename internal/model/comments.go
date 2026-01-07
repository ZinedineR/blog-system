package model

import (
	"blog-system/internal/entity"
	"blog-system/pkg/constant"
	"context"

	"github.com/google/uuid"
)

type BaseCommentsReq struct {
	PostReferencesId string  `json:"post_references_id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserReferencesId *string `json:"user_references_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Content          string  `json:"content" validate:"required" example:"This is a great post! Thanks for sharing."`
}

type CreateCommentReq struct {
	BaseCommentsReq
}

func (req BaseCommentsReq) ToEntity(ctx context.Context) *entity.Comments {
	userReferencesId := constant.GetUserRefId(ctx)
	if userReferencesId != "" {
		req.UserReferencesId = &userReferencesId
	}
	return &entity.Comments{
		ReferencesId:     uuid.NewString(),
		PostReferencesId: req.PostReferencesId,
		UserReferencesId: req.UserReferencesId,
		Content:          req.Content,
	}
}

type CreateCommentRes struct {
	entity.Comments
}

type UpdateCommentReq struct {
	BaseCommentsReq
	ReferencesId string
}
type UpdateCommentRes struct {
	entity.Comments
}

type DeleteCommentReq struct {
	ReferencesId string //uuid, will be get from query handler
}
type DeleteCommentRes struct {
	ReferencesId string //uuid, will be get from query handler
}

type GetAllCommentsReq struct {
	ListReq
}

type GetCommentsByPostReq struct {
	PostReferencesId string `json:"post_references_id"`
}

type GetAllCommentsRes struct {
	PaginationData[entity.Comments]
}

type GetCommentByIDReq struct {
	ReferencesId string `json:"references_id" name:"references_id"`
}

type GetCommentByIDRes struct {
	entity.Comments
}
