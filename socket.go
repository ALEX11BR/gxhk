package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func GetSocket(socketPath string) (socket net.Listener, err error) {
	err = os.RemoveAll(socketPath)
	if err != nil {
		return nil, err
	}

	return net.Listen("unix", socketPath)
}

func HandleConnection(conn net.Conn) {
	decoder := json.NewDecoder(conn)
	var command Args
	decoder.Decode(&command)

	encoder := json.NewEncoder(conn)
	res := NoError

	switch {
	case command.Bind != nil:
		err := Bind(*command.Bind)
		if err != nil {
			res.Status = 1
			res.Message = err.Error()
		}
	case command.Unbind != nil:
		err := Unbind(*command.Unbind)
		if err != nil {
			res.Status = 1
			res.Message = err.Error()
		}
	case command.Info != nil:
		hotkey := command.Info.Hotkey
		if hotkey != "" {
			res.Message = GetInfo(command.Info.Hotkey)
		} else {
			res.Message = GetAllInfo()
		}
	}

	encoder.Encode(res)

	conn.Close()
}

func SocketLoop() {
	for {
		conn, err := Socket.Accept()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		go HandleConnection(conn)
	}
}
