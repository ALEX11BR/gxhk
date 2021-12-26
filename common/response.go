package common

// The Response stores the response which gets displayed.
// WARNING: Make sure that yout Message doesn't end in a newline!!!
// That's beacuse the client will display a newline after the Message.
type Response struct {
	Status  int
	Message string
}

var NoError = Response{0, ""}
