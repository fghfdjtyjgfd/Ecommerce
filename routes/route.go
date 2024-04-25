package routes

import (
	"github.com/gin-gonic/gin"

	"ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/user/signup", controllers.Signup())
	incomingRoutes.POST("/user/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/user/viewproduct", controllers.SearchProduct())
	incomingRoutes.GET("/user/search", controllers.SearchProductByQuery())
}
