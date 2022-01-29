package entities

type Menu struct {
	Id     string  `json:"id" bson:"_id"`
	Name   string  `json:"name" bson:"name"`
	Image  string  `json:"image_url" bson:"image_url"`
	Cost   float32 `json:"cost" bson:"cost"`
	Rating int8    `json:"rating" bson:"rating"`
}
