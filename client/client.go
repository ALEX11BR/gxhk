package client

import (
	"encoding/json"
	"net"

	"github.com/alex11br/gxhk/common"
)

// Just SendCommand, let the daemon do the heavy work, and get through the socket
// the readily formated response that you'll simply return.
func SendCommand(args common.Args) (res common.Response) {
	conn, err := net.Dial("unix", args.SocketPath)
	if err != nil {
		return common.Response{
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
