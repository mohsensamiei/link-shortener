package ginext

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAuthBearerToken(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	dump := strings.Split(token, "Bearer")
	if len(dump) != 2 {
		return ""
	}
	return strings.TrimSpace(dump[1])
}
