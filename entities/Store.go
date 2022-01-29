package entities

type Store struct {
	Id       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Image    string  `json:"image_url" bson:"image_url"`
	Location string  `json:"location" bson:"location"`
	Lat      float32 `json:"lat" bson:"lat"`
	Long     float32 `json:"long" bson:"long"`
}
