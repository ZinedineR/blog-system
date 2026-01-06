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

func (h UserHTTPHandler) TestToken(ctx *gin.Context) {
	request := model.GetUserByIDReq{}
	res, errException := h.UserService.Detail(ctx, &request)
	if errException != nil {
		h.ExceptionJSON(ctx, errException)
		return
	}
	h.DataJSON(ctx, res)
}
