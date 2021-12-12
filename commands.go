package main

import (
	"os/exec"

	"github.com/jezek/xgbutil/keybind"
)

func Exec(command string) {
	exec.Command("sh", "-c", command).Run()
}

func Bind(bindArgs BindCmd) error {
	mods, keys, err := keybind.ParseString(X, bindArgs.Hotkey)
	if err != nil {
		return err
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

	if bindArgs.Released {
		KeyReleaseDescriptions[bindArgs.Hotkey] = bindArgs.Description
	} else {
		KeyPressDescriptions[bindArgs.Hotkey] = bindArgs.Description
	}

	return nil
}

func Unbind(unbindArgs UnbindCmd) error {
	mods, keys, err := keybind.ParseString(X, unbindArgs.Hotkey)
	if err != nil {
		return err
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
		delete(KeyReleaseDescriptions, unbindArgs.Hotkey)
	} else {
		delete(KeyPressDescriptions, unbindArgs.Hotkey)
	}

	return nil
}
