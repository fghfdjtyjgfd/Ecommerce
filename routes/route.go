package routes

import (
	"github.com/gofiber/fiber/v2"

	"ecommerce/controllers"
)

func NewRouter(c *fiber.Ctx) {
	app := fiber.New()

	app.Post("/user/signup", controllers.Signup)
	app.Post("/user/login", controllers.Login)
	app.Post("/admin/addproduct", controllers.ProductViewerAdmin)
	app.Get("/user/viewproduct", controllers.SearchProduct)
	app.Get("/user/search", controllers.SearchProductByQuery)
}
