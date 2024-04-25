package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_name      *string            `json:"first_name"		validate:"required,min=2,max=30"`
	Last_name       *string            `json:"last_name"		validate:"required,min=2,max=30"`
	Password        *string            `json:"password"			validate:"required,min=6"`
	Email           *string            `json:"email"			validate:"required"`
	Phone           *string            `json:"phone"			validate:"required"`
	Token           *string            `json:"token"`
	Refresh_token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	UserID          string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_details []Address          `json:"address" bson:"address"`
	Order_status    []Order            `json:"order" bson:"order"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_name *string            `json:"product_name"`
	Price        *uint64            `json:"price"`
	Rating       *uint8             `json:"rating"`
	Image        *string            `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `bson:"_id"`
	Order_cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_at     time.Time          `json:"oredered_at" bson:"oredered_at"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool
	COD     bool
}
