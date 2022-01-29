package usecases

import (
	"meeting-service/entities"
	"meeting-service/usecases/in"
	"meeting-service/usecases/out"
)

type retrieveMeetingDetailsInteractor struct {
	meetingDataSource out.MeetingDataSource
}

func NewRetrieveMeetingDetailsInteractor(
	meetingDataSource *out.MeetingDataSource,
) in.RetrieveMeetingDetailsInput {
	return &retrieveMeetingDetailsInteractor{
		*meetingDataSource,
	}
}

func (r retrieveMeetingDetailsInteractor) Get(
	meetingId string,
) (response entities.Meeting) {
	return r.meetingDataSource.Get(meetingId)
}
