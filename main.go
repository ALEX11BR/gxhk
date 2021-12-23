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
		if args.ConfigFiles != nil {
			parser.Fail("You can't specify config files if you want to send commands to the running daemon.")
		}

		res := client.SendCommand(args)
		if res.Message != "" {
			fmt.Fprintln(os.Stderr, res.Message)
		}
		os.Exit(res.Status)
	} else {
		os.Exit(daemon.StartDaemon(args))
	}
}
