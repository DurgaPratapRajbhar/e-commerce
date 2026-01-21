package router

import (
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)
 
func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	client.NewAuthClient()
	
	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})
	return r
}
