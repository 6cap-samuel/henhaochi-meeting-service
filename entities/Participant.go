package entities

type Participant struct {
	Id              string `json:"id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	IsParticipating bool   `json:"is_participating" bson:"is_participating"`
}
