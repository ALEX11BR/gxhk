package main

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var IgnoreMods uint16 = xproto.ModMask2 | xproto.ModMaskLock

type Hotkey struct {
	mods uint16
	key  xproto.Keycode
}

func HandleEvent(event xgb.Event) {
	switch ev := event.(type) {
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

func HandleHotkeys() {
	for {
		event, _ := X.Conn().WaitForEvent()
		go HandleEvent(event)
	}
}
