package route

import (
	"blog-system/internal/delivery/http"
	api "blog-system/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	App            *gin.Engine
	Middleware     *api.Middleware
	UserHandler    *http.UserHTTPHandler
	PostHandler    *http.PostHTTPHandler
	CommentHandler *http.CommentHTTPHandler
}

func (h *Router) Setup() {
	h.App.Use(h.Middleware.ErrorHandler)
	baseApi := h.App.Group("")
	baseApi.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	{
		userApi := baseApi.Group("/users")
		{
			userApi.POST("/register", h.UserHandler.Register)
			userApi.POST("/login", h.UserHandler.Login)
			userApi.GET("/try", h.Middleware.JWTAuthentication, h.UserHandler.TestToken)
		}
	}
	{
		postApi := baseApi.Group("/posts")
		postApi.Use(h.Middleware.JWTAuthentication)
		{
			postApi.POST("", h.PostHandler.Create)
			postApi.GET("", h.PostHandler.Find)
			postApi.GET("/:id", h.PostHandler.Detail)
			postApi.PATCH("/:id", h.PostHandler.Update)
			postApi.DELETE("/:id", h.PostHandler.Delete)
		}
	}
	{
		commentApi := baseApi.Group("/comments")
		commentApi.Use(h.Middleware.OptionalJWTAuthentication)
		{
			commentApi.POST("", h.CommentHandler.Create)
			commentApi.GET("", h.CommentHandler.Find)
			commentApi.GET("/:id", h.CommentHandler.Detail)
			commentApi.PATCH("/:id", h.CommentHandler.Update)
			commentApi.DELETE("/:id", h.CommentHandler.Delete)
		}
	}
}
