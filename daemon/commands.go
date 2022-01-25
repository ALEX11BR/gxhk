package daemon

import (
	"fmt"
	"os/exec"

	"github.com/alex11br/gxhk/common"
)

// Exec executes the given sh command string. This method is blocking.
func Exec(command string) {
	exec.Command("sh", "-c", command).Run()
}

// MakeDescription handles the logic behind generating a dummy description for the commands
// which don't come with a description (i.e. it is empty) from the command string,
// or returning the description if it isn't empty
func MakeDescription(description string, command string) string {
	if description != "" {
		return description
	} else {
		return fmt.Sprintf("Spawn '%s'", command)
	}
}

// AddEventInfo handles the logic behind creating a new line and filling it
// with the infos necessary for the 'info' command for each bound hotkey
func AddEventInfo(infos *string, hotkey string, eventName string, description string) {
	*infos += fmt.Sprintf("On %s %s: %s\n", hotkey, eventName, description)
}

// Bind does the heavy work of binding a hotkey (plus its eventual sisters) to a command.
func Bind(bindArgs common.BindCmd) error {
	// First, we find the Hotkey's that correspond to the given hotkey string.
	hotkeys, err := HotkeysFromStr(bindArgs.Hotkey)
	if err != nil {
		return err
	}

	// Then, grab the matching keys which haven't been grabbed yet (i.e. those without commands).
	// If any of these new keys fail to grab, we ungrab what we got to grab before and leave.
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

	// Now we'll set all the appropiate values in the necessary places
	// to successfully bind the hotkeys.
	// It is worth noting that if there are multiple hotkeys for the same string,
	// in the 'info' command only one of them should provide descriptions. As such:
	// CONVENTION: Empty description of a Hotkey (description == "") MEANS ignore that Hotkey!!!
	// We'll put a non-empty description for the first matching Hotkey (refer to MakeDescription)
	// and for any eventual ones an empty description
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

// Unbind does the not-so-heavy work of unbinding a hotkey (plus its eventual sisters),
// which consists in ungrabbing the keys, and deleting the keys in the HotkeyMap's
func Unbind(unbindArgs common.UnbindCmd) error {
	hotkeys, err := HotkeysFromStr(unbindArgs.Hotkey)
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

// GetInfo returns the information about what is bound to the given hotkeyStr.
// It returns a Response structure to easily accomodate the need to mention
// parsing errors if the hotkeyStr is "naughty".
// The first Hotkey to be parsed should be the main one,
// which represents our needed description in the description maps.
func GetInfo(hotkeyStr string) (res common.Response) {
	hotkeys, err := HotkeysFromStr(hotkeyStr)
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

// GetAllInfo creates a string with all the bound hotkeys there are,
// both those bound on press and those bound on release.
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
