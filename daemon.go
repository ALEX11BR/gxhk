package main

import (
	"log"
	"net"
	"os/exec"

	"github.com/jezek/xgbutil"
	"github.com/jezek/xgbutil/keybind"
)

var (
	Socket net.Listener
	X      *xgbutil.XUtil

	KeyPressCommands     = NewCommandsMap()
	KeyPressDescriptions = NewDescriptionsMap()

	KeyReleaseCommands     = NewCommandsMap()
	KeyReleaseDescriptions = NewDescriptionsMap()
)

func StartDaemon(args Args) {
	var err error
	Socket, err = GetSocket(args.SocketPath)
	if err != nil {
		log.Fatal(err)
	}
	defer Socket.Close()

	X, err = xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	keybind.Initialize(X)

	go SocketLoop()

	configFiles := append(args.ConfigFiles, "/etc/gxhkrc")
	for _, configFile := range configFiles {
		go exec.Command(configFile).Run()
	}

	HandleHotkeys()
}
