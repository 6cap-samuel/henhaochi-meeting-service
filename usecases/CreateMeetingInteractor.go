package usecases

import (
	"meeting-service/entities"
	"meeting-service/usecases/in"
	"meeting-service/usecases/out"
)

type createMeetingInput struct {
	meetingDataSource out.MeetingDataSource
}

func NewCreateMeetingInput(
	meeting *out.MeetingDataSource,
) in.CreateMeetingInput {
	return &createMeetingInput{
		*meeting,
	}
}

func (c createMeetingInput) Create(
	meeting entities.Meeting,
) {
	c.meetingDataSource.Create(meeting)
}
