package http

import (
	_ "blog-system/internal/delivery/http/response"
	"blog-system/internal/model"
	service "blog-system/internal/services"

	"github.com/gin-gonic/gin"
)

type CommentHTTPHandler struct {
	Handler
	CommentService service.CommentService
}

func NewCommentHTTPHandler(example service.CommentService) *CommentHTTPHandler {
	return &CommentHTTPHandler{
		CommentService: example,
	}
}

func (h CommentHTTPHandler) Create(ctx *gin.Context) {
	request := model.CreateCommentReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.CommentService.Create(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h CommentHTTPHandler) Find(ctx *gin.Context) {
	var req model.GetAllCommentsReq
	var err error
	req.Page, req.Order, req.Filter, err = h.ParsePaginationParams(ctx)
	if err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	result, errException := h.CommentService.Find(ctx, &req)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, result)
}

func (h CommentHTTPHandler) Detail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.GetCommentByIDReq{
		ReferencesId: idStr,
	}
	res, errException := h.CommentService.Detail(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h CommentHTTPHandler) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.UpdateCommentReq{
		ReferencesId: idStr,
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.CommentService.Update(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

func (h CommentHTTPHandler) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	request := model.DeleteCommentReq{
		ReferencesId: idStr,
	}
	res, errException := h.CommentService.Delete(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}
