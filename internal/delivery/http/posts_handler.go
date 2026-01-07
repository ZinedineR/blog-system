package http

import (
	_ "blog-system/internal/delivery/http/response"
	"blog-system/internal/model"
	service "blog-system/internal/services"

	"github.com/gin-gonic/gin"
)

type PostHTTPHandler struct {
	Handler
	PostService service.PostService
}

func NewPostHTTPHandler(example service.PostService) *PostHTTPHandler {
	return &PostHTTPHandler{
		PostService: example,
	}
}

func (h PostHTTPHandler) Create(ctx *gin.Context) {
	request := model.CreatePostReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.PostService.Create(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h PostHTTPHandler) Find(ctx *gin.Context) {
	var req model.GetAllPostsReq
	var err error
	req.Page, req.Order, req.Filter, err = h.ParsePaginationParams(ctx)
	if err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	result, errException := h.PostService.Find(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

func (h PostHTTPHandler) Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.GetPostByIDReq{
		ReferencesId: idStr,
	}
	res, errException := h.PostService.Detail(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h PostHTTPHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.UpdatePostReq{
		ReferencesId: idStr,
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.PostService.Update(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h PostHTTPHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.DeletePostReq{
		ReferencesId: idStr,
	}
	res, errException := h.PostService.Delete(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}
