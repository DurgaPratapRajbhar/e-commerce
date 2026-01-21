package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/product")
	CategoryRoutes(v1, db)
	ProductRoutes(v1, db)
	ImageRoutes(v1, db)
	FrontendRoutes(v1, db)
	ReviewRoutes(v1, db)
	UnitRoutes(v1, db)
}
