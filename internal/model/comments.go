package model

import (
	"blog-system/internal/entity"
	"blog-system/pkg/constant"
	"context"

	"github.com/google/uuid"
)

type BaseCommentsReq struct {
	PostReferencesId string  `json:"post_references_id" validate:"required"`
	UserReferencesId *string `json:"user_references_id"`
	Content          string  `json:"content" validate:"required"`
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
	BaseCommentsReq
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
