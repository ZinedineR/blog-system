package api

import (
	"blog-system/internal/delivery/http"
	"blog-system/pkg/constant"
	"blog-system/pkg/signature"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	signaturer signature.Signaturer
	http.Handler
}

func NewMiddleware(signaturer signature.Signaturer) *Middleware {
	return &Middleware{
		signaturer: signaturer,
	}
}
func (m *Middleware) ErrorHandler(c *gin.Context) {

	defer func() {
		if err0 := recover(); err0 != nil {
			slog.Any("error", err0)
			m.InternalErrorJSON(c, "Request is halted unexpectedly, please contact the administrator.", err0)
		}
	}()
	c.Next()
}
func (m *Middleware) JWTAuthentication(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		m.UnauthorizedJSON(c, "Invalid token")
		return
	}
	token := authFields[1]
	jwtCheck, err := m.signaturer.JWTCheck(token)
	if err != nil {
		m.ExceptionJSON(c, err)
		return
	}
	constant.SetUserReferencesId(c, &jwtCheck.UserReferencesId)
	c.Next()
}

//AsaDMIN
