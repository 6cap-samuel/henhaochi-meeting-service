package entities

import (
	"time"
)

type Meeting struct {
	Id              string        `json:"id" bson:"_id"`
	Store           Store         `json:"store" bson:"store"`
	MeetingDateTime time.Time     `json:"meeting_date_time" bson:"meeting_date_time"`
	MenusImageUrl   []string      `json:"menus_image_url" bson:"menus_image_url"`
	Participants    []Participant `json:"participants" bson:"participants"`
}
