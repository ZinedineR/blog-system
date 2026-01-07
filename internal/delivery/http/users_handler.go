package http

import (
	_ "blog-system/internal/delivery/http/response"
	"blog-system/internal/model"
	service "blog-system/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHTTPHandler struct {
	Handler
	UserService service.UserService
}

func NewUserHTTPHandler(example service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		UserService: example,
	}
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	Registers a new user in the system
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.CreateUserReq	true	"Register User Request"
//	@Success		200		{object}	response.DataResponse{data=model.CreateUserRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse									"error"
//	@Router			/users/register [post]
func (h UserHTTPHandler) Register(ctx *gin.Context) {
	request := model.CreateUserReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.UserService.Register(ctx, &request)

	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}

// Login godoc
//
//	@Summary		User login
//	@Description	Authenticates a user and returns a token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		model.CreateUserReq	true	"Login User Request"
//	@Success		200		{object}	response.DataResponse{data=model.CreateUserRes}	"success"
//	@Failure		400		{object}	response.ErrorResponse									"error"
//	@Router			/users/login [post]
func (h UserHTTPHandler) Login(ctx *gin.Context) {
	request := model.CreateUserReq{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		h.BadRequestJSON(ctx, err.Error())
		return
	}
	res, errException := h.UserService.Login(ctx, &request)

	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}

	h.DataJSON(ctx, res)
}
