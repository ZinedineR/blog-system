package constant

import (
	"context"

	"github.com/gin-gonic/gin"
)

const (
	UserReferencesId string = "user_references_id"
)

func SetUserReferencesId(c *gin.Context, references *string) {
	c.Set(UserReferencesId, *references)
}

func GetUserRefId(ctx context.Context) string {
	data, ok := ctx.Value(UserReferencesId).(string)
	if ok {
		return data
	}
	return ""
}
