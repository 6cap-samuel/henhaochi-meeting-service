package out

import "meeting-service/entities"

type MeetingDataSource interface {
	Get(meetingId string) (response entities.Meeting)
	Create(meeting entities.Meeting)
}
