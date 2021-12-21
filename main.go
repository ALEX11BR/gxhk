package main

import (
	"fmt"
	"os"
)

func main() {
	args, parser := ParseArgs()

	if parser.Subcommand() != nil {
		if args.ConfigFiles != nil {
			parser.Fail("You can't specify config files if you want to send commands to the running daemon.")
		}

		res := SendCommand(args)
		if res.Message != "" {
			fmt.Fprintln(os.Stderr, res.Message)
		}
		os.Exit(res.Status)
	} else {
		StartDaemon(args)
	}
}
