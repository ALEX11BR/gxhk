package main

import (
	"github.com/jezek/xgb/xproto"
)

var IgnoreMods uint16 = xproto.ModMask2 | xproto.ModMaskLock

type Hotkey struct {
	mods uint16
	key  xproto.Keycode
}

func HandleHotkeys() {
	for {
		ev, _ := X.Conn().WaitForEvent()

		switch ev := ev.(type) {
		case xproto.KeyPressEvent:
			hotkey := Hotkey{
				mods: ev.State | IgnoreMods,
				key:  ev.Detail,
			}
			Exec(KeyPressCommands.Get(hotkey))
		case xproto.KeyReleaseEvent:
			hotkey := Hotkey{
				mods: ev.State | IgnoreMods,
				key:  ev.Detail,
			}
			Exec(KeyReleaseCommands.Get(hotkey))
		}
	}
}
