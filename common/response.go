package common

type Response struct {
	Status  int
	Message string
}

var NoError = Response{0, ""}
