package daemon

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/alex11br/gxhk/common"
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
	var command common.Args
	decoder.Decode(&command)

	encoder := json.NewEncoder(conn)
	res := common.NoError

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
