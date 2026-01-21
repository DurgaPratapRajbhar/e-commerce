package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"
	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func UnitRoutes(route *gin.RouterGroup, db *gorm.DB) {

	proRepo := repository.NewProductUnitRepository(db)
	proService := services.NewProductUnitService(proRepo)

	Controller := controllers.NewProductUnitController(proService)
	UnitGroup := route.Group("/product-units")
	UnitGroup.Use(middleware.ServiceAuthMiddleware())
	{
		UnitGroup.POST("", Controller.CreateUnit)
		UnitGroup.GET("/:id", Controller.GetUnit)
		UnitGroup.GET("", Controller.GetAllUnits)
		UnitGroup.PUT("/:id", Controller.UpdateUnit)
		UnitGroup.DELETE("/:id", Controller.DeleteUnit)
	}

}
