package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"ecommerce/database"
	"ecommerce/controllers"
	"ecommerce/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	apps := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	app := fiber.New()
	app.Use(middleware.Authentication)

	app.Get("/addtocart", app.AddToCart)
	app.Get("/removeitem", app.RemoveItem)
	app.Get("/cartcheckout", app.CartCheckOut)
	app.Get("/instantbuy", app.InstantBuy)

	app.Listen(":" + port)
}
