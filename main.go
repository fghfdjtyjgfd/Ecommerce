package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	router.GET("/viewcart", app.GetItemFromCart())

	router.POST("/addaddress", app.AddAddress())
	router.PUT("/edithomeaddress", app.EditHomeAddress())
	router.PUT("/editworkaddress", app.EditWorkAddress())
	router.DELETE("/deleteaddress", app.DeleteAddress())

	log.Fatal(router.Run(":" + port))
}
