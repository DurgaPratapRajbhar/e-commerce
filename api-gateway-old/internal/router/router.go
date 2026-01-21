// package router

// import (
 
// 	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
// 	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
// 	"github.com/gin-gonic/gin"
// 	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
// )

// func SetupRouter(cfg *config.Config) *gin.Engine {
// 	r := gin.New()
// 	r.Use(gin.Logger())
// 	r.Use(gin.Recovery())
// 	r.Use(middleware.CORSMiddleware())

// 	sp := proxy.NewServiceProxy(cfg)
// 	ap := proxy.NewAggregateProxy(sp)  // CHANGE THIS LINE - pass sp instead of cfg

// 	api := r.Group("/api/v1")
// 	{
// 		// Auth routes (public)
// 		auth := api.Group("/auth")
// 		{
// 			auth.Any("/register", sp.ProxyRequest("auth"))
// 			auth.Any("/login", sp.ProxyRequest("auth"))
// 			auth.Any("/logout", sp.ProxyRequest("auth"))
// 			auth.Any("/refresh", sp.ProxyRequest("auth"))
// 			auth.Any("/password/*path", sp.ProxyRequest("auth"))
// 			auth.Any("/verification/*path", sp.ProxyRequest("auth"))

// 			// Protected auth routes
// 			authProtected := auth.Group("")
// 			authProtected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
// 			{
// 				authProtected.Any("/me", sp.ProxyRequest("auth"))
// 				authProtected.Any("/validate", sp.ProxyRequest("auth"))
// 				authProtected.Any("/permissions/*path", sp.ProxyRequest("auth"))
// 			}
// 		}

// 		// Admin routes (protected)
// 		admin := api.Group("/admin")
// 		admin.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
// 		{
// 			admin.Any("/users", ap.AggregateAdminUsers)
// 		}

// 		// Protected service routes
// 		protected := api.Group("")
// 		protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
// 		{
// 			// User service
// 			protected.Any("/user", sp.ProxyRequest("user"))
// 			protected.Any("/user/*path", sp.ProxyRequest("user"))

// 			// Cart service
// 			protected.Any("/cart", sp.ProxyRequest("cart"))
// 			protected.Any("/cart/*path", sp.ProxyRequest("cart"))

// 			// Order service
// 			protected.Any("/order", sp.ProxyRequest("order"))
// 			protected.Any("/order/*path", sp.ProxyRequest("order"))

// 			// Payment service
// 			protected.Any("/payment", sp.ProxyRequest("payment"))
// 			protected.Any("/payment/*path", sp.ProxyRequest("payment"))

// 			// Product service
// 			protected.Any("/product", sp.ProxyRequest("product"))
// 			protected.Any("/product/*path", sp.ProxyRequest("product"))

// 			// Shipment service
// 			protected.Any("/shipment", sp.ProxyRequest("shipment"))
// 			protected.Any("/shipment/*path", sp.ProxyRequest("shipment"))

// 			// Inventory service
// 			protected.Any("/inventory", sp.ProxyRequest("inventory"))
// 			protected.Any("/inventory/*path", sp.ProxyRequest("inventory"))
// 		}
// 	}

// 	// Health check
// 	r.GET("/health", sp.HealthCheck)

// 	return r
// }





package router

import (
 
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	sp := proxy.NewServiceProxy(cfg)
	ap := proxy.NewAggregateProxy(sp)

	api := r.Group("/api/v1")
	{
		// Set up different route groups
		SetupAuthRoutes(api, sp, cfg)
		SetupAdminRoutes(api, sp, ap, cfg)
		SetupUserRoutes(api, sp, cfg)
		SetupProductRoutes(api, sp, cfg)
		SetupCartRoutes(api, sp, cfg)
		SetupOrderRoutes(api, sp, cfg)
		SetupPaymentRoutes(api, sp, cfg)
		SetupShippingRoutes(api, sp, cfg)
		SetupInventoryRoutes(api, sp, cfg)
		SetupUserListRoutes(api, cfg)
				SetupAggregatorRoutes(api, cfg)
	}

	// Health check
	r.GET("/health", sp.HealthCheck)

	return r
}