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

// Create godoc
//
//	 @Security BearerAuth
//		@Summary		Create a new post
//		@Description	Creates a new post in the system
//	 @Security BearerAuth
//		@Tags			Post
//		@Accept			json
//		@Produce		json
//		@Param			post	body		model.CreatePostReq true	"Create Post Request"
//		@Success		200		{object}	response.DataResponse{data=model.CreatePostRes}	"success"
//		@Failure		400		{object}	response.ErrorResponse										"error"
//		@Router			/posts [post]
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

// Find godoc
//
//	@Summary		Find posts
//	@Description	Finds posts in the system with a paginated list of vendors with optional filters and sorting
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//
// @Param offset query int false "Offset for pagination"
// @Param limit query int false "Limit for pagination"
// @Param filter query string false "Filter rules<br><br>### Format:<br>{field}:{value}:{operator}<br>Supported operators: eq, lt, gt, lte, gte, in, like, is, not"
// @Param order query string false "Sort rules<br><br>### Format:<br>{field}:{direction}<br>Supported directions: asc, desc"
//
// @Success 200 {object} response.PaginationResponse{data=[]model.GetPostByIDRes} "success"
// @Failure 400 {object} response.ErrorResponse "error"									"error"
//
//	@Router			/posts [get]
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

// Detail godoc
//
//	@Summary		Get a post
//	@Description	Gets a post in the system
//
// @Param id path string true "Post ID"
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.DataResponse{data=model.GetPostByIDRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/posts/{id} [get]
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

// Update godoc
//
//	 @Security BearerAuth
//		@Summary		Update a post
//		@Description	Updates a post in the system
//
// @Param id path string true "Post ID"
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Param			post	body		model.UpdatePostReq true	"Update Post Request"
//	@Success		200		{object}	response.DataResponse{data=model.UpdatePostRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/posts/{id} [patch]
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

// Delete godoc
//
//	 @Security BearerAuth
//		@Summary		Delete a post
//		@Description	Deletes a post in the system
//
// @Param id path string true "Post ID"
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.DataResponse{data=model.DeletePostRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse										"error"
//	@Router			/posts/{id} [delete]
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
