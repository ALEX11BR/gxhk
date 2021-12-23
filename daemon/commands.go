package daemon

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/alex11br/gxhk/common"
	"github.com/jezek/xgbutil/keybind"
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
	mods, keys, err := keybind.ParseString(X, bindArgs.Hotkey)
	if err != nil {
		return err
	}

	if mods&IgnoreMods > 0 {
		return errors.New("hotkeys can't rely on the status of the Caps Lock or the Num Lock (Mod2)")
	}

	for _, key := range keys {
		hotkey := Hotkey{
			mods | IgnoreMods,
			key,
		}

		err = keybind.GrabChecked(X, X.RootWin(), mods, key)
		if err != nil {
			return err
		}

		if bindArgs.Released {
			KeyReleaseCommands.Set(hotkey, bindArgs.RunCommand)
		} else {
			KeyPressCommands.Set(hotkey, bindArgs.RunCommand)
		}
	}

	description := MakeDescription(bindArgs.Description, bindArgs.RunCommand)
	if bindArgs.Released {
		KeyReleaseDescriptions.Set(bindArgs.Hotkey, description)
	} else {
		KeyPressDescriptions.Set(bindArgs.Hotkey, description)
	}

	return nil
}

func Unbind(unbindArgs common.UnbindCmd) error {
	mods, keys, err := keybind.ParseString(X, unbindArgs.Hotkey)
	if err != nil {
		return err
	}

	if mods&IgnoreMods > 0 {
		return errors.New("hotkeys can't rely on the status of the Caps Lock or the Num Lock (Mod2)")
	}

	for _, key := range keys {
		hotkey := Hotkey{
			mods | IgnoreMods,
			key,
		}

		if unbindArgs.Released {
			KeyReleaseCommands.Delete(hotkey)
		} else {
			KeyPressCommands.Delete(hotkey)
		}

		if KeyPressCommands.IsEmpty(hotkey) && KeyReleaseCommands.IsEmpty(hotkey) {
			keybind.Ungrab(X, X.RootWin(), mods, key)
		}
	}

	if unbindArgs.Released {
		KeyReleaseDescriptions.Delete(unbindArgs.Hotkey)
	} else {
		KeyPressDescriptions.Delete(unbindArgs.Hotkey)
	}

	return nil
}

func GetInfo(hotkey string) (info string) {
	pressInfo := KeyPressDescriptions.Get(hotkey)
	if pressInfo != "" {
		AddEventInfo(&info, hotkey, "press", pressInfo)
	}

	releaseInfo := KeyReleaseDescriptions.Get(hotkey)
	if releaseInfo != "" {
		AddEventInfo(&info, hotkey, "release", releaseInfo)
	}

	return
}

func GetAllInfo() (info string) {
	KeyPressDescriptions.Iter(func(hotkey, description string) {
		AddEventInfo(&info, hotkey, "press", description)
	})

	KeyReleaseDescriptions.Iter(func(hotkey, description string) {
		AddEventInfo(&info, hotkey, "release", description)
	})

	return
}
