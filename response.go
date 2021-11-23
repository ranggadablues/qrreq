package qrreq

import "time"

type IsUserAllowedToLoginResponse struct {
	Status    bool
	Token     string
	Mobile    string
	Email     string
	Name      string
	LastLogin time.Time
}
