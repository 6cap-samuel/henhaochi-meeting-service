package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"meeting-service/configurations"
	"meeting-service/entities"
	"meeting-service/exceptions"
	"meeting-service/usecases/out"
)

type meetingRepository struct {
	Collection *mongo.Collection
}

func NewMeetingRepository(
	database *mongo.Database,
) out.MeetingDataSource {
	return &meetingRepository{
		Collection: database.Collection("meetings"),
	}
}

func (m meetingRepository) Get(
	meetingId string,
) (response entities.Meeting) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	err := m.Collection.FindOne(
		ctx,
		bson.D{
			{"_id",
				meetingId,
			},
		},
	).Decode(&response)

	exceptions.PanicIfNeeded(err)

	return response
}

func (m meetingRepository) Create(
	meeting entities.Meeting,
) {
	ctx, cancel := configurations.NewMongoContext()
	defer cancel()

	_, err := m.Collection.InsertOne(ctx, meeting)

	exceptions.PanicIfNeeded(err)
}
