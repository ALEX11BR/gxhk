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

	KeyPressCommands     = NewRWLockedMap()
	KeyPressDescriptions = make(map[string]string)

	KeyReleaseCommands     = NewRWLockedMap()
	KeyReleaseDescriptions = make(map[string]string)
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
		command := exec.Command(configFile)
		command.Start()
	}

	HandleHotkeys()
}
