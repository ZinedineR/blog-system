package model

import (
	"blog-system/internal/entity"
	"blog-system/pkg/constant"
	"context"
	"time"

	"github.com/google/uuid"
)

type BasePostsReq struct {
	Title           string     `json:"title" validate:"required" example:"My First Blog Post"`
	Content         string     `json:"content" validate:"required" example:"This is the content of my blog post"`
	Tags            string     `json:"tags" example:"golang,programming,tutorial"`
	Published       bool       `json:"published" example:"true"`
	LastPublishedAt *time.Time `json:"-"`
}

type CreatePostReq struct {
	BasePostsReq
}

func (req BasePostsReq) ToEntity(ctx context.Context) *entity.Posts {
	if req.Published {
		now := time.Now()
		req.LastPublishedAt = &now
	}
	return &entity.Posts{
		ReferencesId:     uuid.NewString(),
		UserReferencesId: constant.GetUserRefId(ctx),
		Title:            req.Title,
		Content:          req.Content,
		Tags:             req.Tags,
		Published:        req.Published,
		LastPublishedAt:  req.LastPublishedAt,
	}
}

type CreatePostRes struct {
	entity.Posts
}

type UpdatePostReq struct {
	BasePostsReq
	ReferencesId string
}
type UpdatePostRes struct {
	entity.Posts
}

type DeletePostReq struct {
	ReferencesId string //uuid, will be get from query handler
}
type DeletePostRes struct {
	ReferencesId string //uuid, will be get from query handler
}

type GetAllPostsReq struct {
	ListReq
}
type GetAllPostsRes struct {
	PaginationData[entity.Posts]
}

type GetPostByIDReq struct {
	ReferencesId string `json:"references_id" name:"references_id"`
}

type GetPostByIDRes struct {
	entity.Posts
}
