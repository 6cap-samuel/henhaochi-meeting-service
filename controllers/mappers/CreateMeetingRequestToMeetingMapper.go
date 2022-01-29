package mappers

import (
	"meeting-service/configurations"
	"meeting-service/controllers/requests"
	"meeting-service/entities"
)

func CreatePostRequestToPostMapper(request requests.CreateMeetingRequest) *entities.Meeting {
	var participants []entities.Participant

	for _, participant := range request.Participants {
		participants = append(participants, entities.Participant{
			Id:              configurations.NewIdentity(),
			Name:            participant,
			IsParticipating: false,
		})
	}

	return &entities.Meeting{
		Id: configurations.NewIdentity(),
		Store: entities.Store{
			Id:       configurations.NewIdentity(),
			Name:     request.StoreName,
			Image:    request.StoreImage,
			Location: request.StoreLocation,
			Lat:      request.StoreLat,
			Long:     request.StoreLong,
		},
		MenusImageUrl:   request.StoreMenuImages,
		MeetingDateTime: request.MeetingDateTime,
		Participants:    participants,
	}
}
