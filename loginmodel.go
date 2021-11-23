package qrreq

import (
	"strings"
)

type LoginRequest struct {
	Username string
	Password string
	Broker   string
	Deviceid string
}

// ValidateInput check qrauth parameter
// For now only check for empty string
// return false if any of username, password, or broker name is empty
// return true otherwise
func (l *LoginRequest) ValidateInput() bool {
	if l.Username == "" ||
		l.Password == "" ||
		l.Broker == "" ||
		strings.Contains(l.Username, "\n") ||
		strings.Contains(l.Password, "\n") ||
		strings.Contains(l.Broker, "\n") ||
		strings.Contains(l.Deviceid, "\n") {
		return false
	}

	return true
}
