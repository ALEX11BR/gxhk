package main

import (
	"encoding/json"
	"net"
)

func SendCommand(args Args) (res Response) {
	conn, err := net.Dial("unix", args.SocketPath)
	if err != nil {
		return Response{
			Status:  1,
			Message: err.Error(),
		}
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	encoder.Encode(args)

	decoder := json.NewDecoder(conn)
	decoder.Decode(&res)
	return
}
