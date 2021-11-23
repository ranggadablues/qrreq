package qrreq

type ForgotPassRequest struct {
	Username string
	Broker   string
	Deviceid string
}

type ForgotPassResponse struct {
	Username string
	Email    string
	Success  bool
}
