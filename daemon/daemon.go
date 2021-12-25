package daemon

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path"

	"github.com/alex11br/gxhk/common"
	"github.com/alex11br/xgbutil"
	"github.com/alex11br/xgbutil/keybind"
)

var (
	Socket net.Listener
	X      *xgbutil.XUtil

	KeyPressCommands     = NewCommandsMap()
	KeyPressDescriptions = NewDescriptionsMap()

	KeyReleaseCommands     = NewCommandsMap()
	KeyReleaseDescriptions = NewDescriptionsMap()
)

func StartDaemon(args common.Args) int {
	var err error
	Socket, err = GetSocket(args.SocketPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer Socket.Close()

	X, err = xgbutil.NewConn()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	keybind.Initialize(X)

	go SocketLoop()

	if !args.NoDefaultConfigs {
		homeConfig := os.Getenv("XDG_CONFIG_HOME")
		if homeConfig == "" {
			homeDir, _ := os.UserHomeDir()
			homeConfig = path.Join(homeDir, ".config")
		}
		homeConfig = path.Join(homeConfig, "gxhk", "gxhkrc")

		go exec.Command("/etc/gxhkrc").Run()
		go exec.Command(homeConfig).Run()
	}
	for _, configFile := range args.ConfigFiles {
		go exec.Command(configFile).Run()
	}

	HandleHotkeys()

	return 0
}
