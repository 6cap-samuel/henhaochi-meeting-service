package in

import "meeting-service/entities"

type CreateMeetingInput interface {
	Create(meeting entities.Meeting)
}
