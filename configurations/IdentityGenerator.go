package configurations

import "github.com/google/uuid"

func NewIdentity() string {
	return uuid.New().String()
}
