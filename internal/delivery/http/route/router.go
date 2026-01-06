package route

import (
	"blog-system/internal/delivery/http"
	api "blog-system/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	App         *gin.Engine
	Middleware  *api.Middleware
	UserHandler *http.UserHTTPHandler
}

func (h *Router) Setup() {
	h.App.Use(h.Middleware.ErrorHandler)
	baseApi := h.App.Group("")
	baseApi.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	{
		userApi := baseApi.Group("/user")
		{
			userApi.POST("/register", h.UserHandler.Register)
			userApi.POST("/login", h.UserHandler.Login)
			userApi.GET("/try", h.Middleware.JWTAuthentication, h.UserHandler.TestToken)
		}
	}
}
