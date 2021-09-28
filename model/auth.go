package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// AuthEmail ...
type AuthEmail struct {
	Email string `json:"email"`
}

//AuthLoginDTO ...
type AuthLoginDTO struct {
	Email          string `validate:"required" json:"email"`
	MasterPassword string `validate:"required" json:"master_password"`
}

//AuthLoginResponse ...
type AuthLoginResponse struct {
	Type            string `json:"type"`
	TransmissionKey string `json:"transmission_key"`
	*UserDTO
	*SubscriptionAuthDTO
}

//TokenDetailsDTO ...
type TokenDetailsDTO struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	AtExpiresTime   time.Time
	RtExpiresTime   time.Time
	AtUUID          uuid.UUID
	RtUUID          uuid.UUID
	TransmissionKey string `json:"transmission_key"`
}
