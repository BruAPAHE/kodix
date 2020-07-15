package models

type AutoAttributes struct {
	Id      int    `json:"id",bson:"auto_id"`
	Brand   string `json:"brand",bson:"brand"`
	Model   string `json:"model",bson:"model"`
	Price   uint   `json:"price",bson:"price"`
	Status  string `json:"status",bson:"status"`
	Mileage int    `json:"mileage",bson:"mileage"`
}
