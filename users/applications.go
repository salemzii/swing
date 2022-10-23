package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/salemzii/swing/logs"
)

type Application struct {
	Id       uuid.UUID        `json:"id"`
	Name     string           `json:"name"`
	AppToken TokenDetails     `json:"apptoken"`
	UserId   uuid.UUID        `json:"userid"`
	Active   bool             `json:"active"`
	Created  time.Time        `json:"created"`
	Updated  time.Time        `json:"updated"`
	Records  []logs.LogRecord `json:"records"`
}

type TokenDetails struct {
	Id         uuid.UUID `json:"id"`
	Token      string    `json:"token"`
	Expires_at time.Time `json:"expires"`
	Rate_limit int       `json:"rate_limit"`
	Enabled    bool      `json:"enabled"`
	Created    time.Time `json:"created"`
}
