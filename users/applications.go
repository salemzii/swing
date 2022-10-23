package users

import (
	"time"
)

type TokenDetails struct {
	Id         int       `json:"id"`
	Token      string    `json:"token"`
	Expires_at time.Time `json:"expires"`
	Rate_limit int       `json:"rate_limit"`
	Enabled    bool      `json:"enabled"`
	Created    time.Time `json:"created"`
	UserId     int       `json:"userid"`
}
