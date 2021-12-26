package main

import (
	"fmt"
	"os"

	"github.com/alex11br/gxhk/client"
	"github.com/alex11br/gxhk/common"
	"github.com/alex11br/gxhk/daemon"
)

func main() {
	args, parser := common.ParseArgs()

	if parser.Subcommand() != nil {
		var stream *os.File
		res := client.SendCommand(args)

		if res.Status == 0 {
			stream = os.Stdout
		} else {
			stream = os.Stderr
		}

		if res.Message != "" {
			fmt.Fprintln(stream, res.Message)
		}
		os.Exit(res.Status)
	} else {
		os.Exit(daemon.StartDaemon(args))
	}
}
