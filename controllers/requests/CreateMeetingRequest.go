package requests

import (
	"time"
)

type CreateMeetingRequest struct {
	StoreName       string    `json:"store_name"`
	StoreImage      string    `json:"store_image"`
	StoreLocation   string    `json:"store_location"`
	StoreLat        float32   `json:"store_lat"`
	StoreLong       float32   `json:"store_long"`
	StoreMenuImages []string  `json:"store_meeting_images"`
	MeetingDateTime time.Time `json:"meeting_date_time"`
	Participants    []string  `json:"participants"`
}
