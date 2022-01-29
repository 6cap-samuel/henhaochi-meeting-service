package in

import "meeting-service/entities"

type RetrieveMeetingDetailsInput interface {
	Get(meetingId string) (response entities.Meeting)
}
