package main

import (
	"encoding/json"
	"log"
	"net"
)

func SendCommand(args Args) (res Response) {
	conn, err := net.Dial("unix", args.SocketPath)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	encoder.Encode(args)

	decoder := json.NewDecoder(conn)
	decoder.Decode(&res)
	return
}
