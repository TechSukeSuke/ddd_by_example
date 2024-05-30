package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k-shimizu04/ddd_by_example/config"
)

type CorsMiddleware struct {
	APIHost string
	APIPort string
}

func NewCorsMiddleware(c *config.APIServer) *CorsMiddleware {
	return &CorsMiddleware{
		APIHost: c.APIHost,
		APIPort: c.APIPort,
	}
}

func (c *CorsMiddleware) Execute(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	// ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Cookie")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS")

	if ctx.Request.Method == "OPTION" {
		ctx.Writer.WriteHeader(http.StatusOK)
		return
	}

	ctx.Next()
}
