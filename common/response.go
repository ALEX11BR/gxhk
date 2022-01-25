package common

// The Response stores the response which gets displayed.
type Response struct {
	Status  int
	Message string
}

var NoError = Response{0, ""}
