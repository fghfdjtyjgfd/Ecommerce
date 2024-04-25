package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"ecommerce/models"
)

func Hashpassword (password string) string {

}

func VerifyPassword (userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(),time.Second*100)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		}

		validateErr := Validate.Struct(user)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr})
			return
		}

		count, err := UserCollection.CountDocument(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return 
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
		}

		count, err = UserCollection.CountDocument(ctx, bson.M{"phone": user.Phone})
		
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return 
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "phone number already in use"})
		}
	
	}
}

func Login() gin.HandlerFunc {
	
}


func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {

}

func SearchProductByQuery() gin.HandlerFunc {
	
}