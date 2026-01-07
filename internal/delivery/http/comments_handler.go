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

// Create godoc
//
//	 @Security BearerAuth
//		@Summary		Create a new comment
//		@Description	Creates a new comment in the system
//		@Tags			Comment
//		@Accept			json
//		@Produce		json
//		@Param			comment	body		model.CreateCommentReq true	"Create Comment Request"
//		@Success		200		{object}	response.DataResponse{data=model.CreateCommentRes}	"success"
//		@Failure		400		{object}	response.ErrorResponse										"error"
//		@Router			/comments [post]
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

// Find godoc
//
//	@Summary		Find comments
//	@Description	Finds comments in the system with a paginated list with optional filters and sorting
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter query string false "Filter rules<br><br>### Format:<br>{field}:{value}:{operator}<br>Supported operators: eq, lt, gt, lte, gte, in, like, is, not"
// @Param order query string false "Sort rules<br><br>### Format:<br>{field}:{direction}<br>Supported directions: asc, desc"
//
// @Success 200 {object} response.PaginationResponse{data=[]model.GetCommentByIDRes} "success"
// @Failure 400 {object} response.ErrorResponse "error"
//
//	@Router			/comments [get]
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

// Detail godoc
//
//	@Summary		Get a comment
//	@Description	Gets a comment in the system
//
// @Param id path string true "Comment ID"
//
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.DataResponse{data=model.GetCommentByIDRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/comments/{id} [get]
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

// Update godoc
//
//	 @Security BearerAuth
//		@Summary		Update a comment
//		@Description	Updates a comment in the system
//
// @Param id path string true "Comment ID"
//
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			comment	body		model.UpdateCommentReq true	"Update Comment Request"
//	@Success		200		{object}	response.DataResponse{data=model.UpdateCommentRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/comments/{id} [patch]
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

// Delete godoc
//
//	 @Security BearerAuth
//		@Summary		Delete a comment
//		@Description	Deletes a comment in the system
//
// @Param id path string true "Comment ID"
//
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.DataResponse{data=model.DeleteCommentRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/comments/{id} [delete]
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
