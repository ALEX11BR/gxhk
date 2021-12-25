package daemon

import (
	"fmt"
	"os/exec"

	"github.com/alex11br/gxhk/common"
)

func Exec(command string) {
	exec.Command("sh", "-c", command).Run()
}

func MakeDescription(description string, command string) string {
	if description != "" {
		return description
	} else {
		return fmt.Sprintf("Spawn '%s'", command)
	}
}

func AddEventInfo(infos *string, hotkey string, eventName string, description string) {
	if *infos != "" {
		*infos += "\n"
	}
	*infos += fmt.Sprintf("On %s %s: %s", hotkey, eventName, description)
}

func Bind(bindArgs common.BindCmd) error {
	hotkeys, err := HotkeyFromStr(bindArgs.Hotkey)
	if err != nil {
		return err
	}

	newlyBound := make([]Hotkey, 0)
	for _, hotkey := range hotkeys {
		if KeyPressCommands.IsEmpty(hotkey) && KeyReleaseCommands.IsEmpty(hotkey) {
			err = hotkey.Grab()
			if err != nil {
				for _, toUnbind := range newlyBound {
					toUnbind.Ungrab()
				}
				return err
			}

			newlyBound = append(newlyBound, hotkey)
		}
	}

	for i, hotkey := range hotkeys {
		var description string
		if i == 0 {
			description = MakeDescription(bindArgs.Description, bindArgs.RunCommand)
		} else {
			description = ""
		}

		if bindArgs.Released {
			KeyReleaseCommands.Set(hotkey, bindArgs.RunCommand)
			KeyReleaseDescriptions.Set(hotkey, description)
		} else {
			KeyPressCommands.Set(hotkey, bindArgs.RunCommand)
			KeyPressDescriptions.Set(hotkey, description)
		}
	}

	return nil
}

func Unbind(unbindArgs common.UnbindCmd) error {
	hotkeys, err := HotkeyFromStr(unbindArgs.Hotkey)
	if err != nil {
		return err
	}

	for _, hotkey := range hotkeys {
		if KeyPressCommands.IsEmpty(hotkey) && KeyReleaseCommands.IsEmpty(hotkey) {
			hotkey.Ungrab()
		}

		if unbindArgs.Released {
			KeyReleaseCommands.Delete(hotkey)
			KeyReleaseDescriptions.Delete(hotkey)
		} else {
			KeyPressCommands.Delete(hotkey)
			KeyPressDescriptions.Delete(hotkey)
		}
	}

	return nil
}

func GetInfo(hotkeyStr string) (res common.Response) {
	hotkeys, err := HotkeyFromStr(hotkeyStr)
	if err != nil {
		return common.Response{
			Status:  1,
			Message: err.Error(),
		}
	}
	hotkey := hotkeys[0]

	pressInfo := KeyPressDescriptions.Get(hotkey)
	if pressInfo != "" {
		AddEventInfo(&res.Message, hotkeyStr, "press", pressInfo)
	}

	releaseInfo := KeyReleaseDescriptions.Get(hotkey)
	if releaseInfo != "" {
		AddEventInfo(&res.Message, hotkeyStr, "release", releaseInfo)
	}

	return
}

func GetAllInfo() (info string) {
	KeyPressDescriptions.ForEach(func(hotkey Hotkey, description string) {
		if description != "" {
			AddEventInfo(&info, hotkey.ToStr(), "press", description)
		}
	})

	KeyReleaseDescriptions.ForEach(func(hotkey Hotkey, description string) {
		if description != "" {
			AddEventInfo(&info, hotkey.ToStr(), "release", description)
		}
	})

	return
}
