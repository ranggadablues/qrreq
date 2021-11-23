package qrreq

import (
	"errors"
)

var (
	ErrorBadRequest     = errors.New("Bad Request")
	ErrorInternalServer = errors.New("Internal Server Error")
	ErrorRemoteServer   = errors.New("Remote Server Error")
)
