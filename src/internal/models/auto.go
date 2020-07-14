package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AutoAttributes struct {
	Id      primitive.ObjectID `bson:"_id"`
	Brand   string             `bson:"brand"`
	Model   string             `bson:"model"`
	Price   uint               `bson:"price"`
	Status  string             `bson:"status"`
	Mileage int                `bson:"mileage"`
}
