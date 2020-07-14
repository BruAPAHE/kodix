package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AutoAttributes struct {
	Id      primitive.ObjectID `json:"id",bson:"_id"`
	Brand   string             `json:"brand",bson:"brand"`
	Model   string             `json:"model",bson:"model"`
	Price   uint               `json:"price",bson:"price"`
	Status  string             `json:"status",bson:"status"`
	Mileage int                `json:"mileage",bson:"mileage"`
}
