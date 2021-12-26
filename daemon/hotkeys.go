package daemon

import (
	"errors"

	"github.com/alex11br/xgbutil/keybind"
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

var IgnoreMods uint16 = xproto.ModMask2 | xproto.ModMaskLock

type Hotkey struct {
	mods uint16
	key  xproto.Keycode
}

// HotkeysFromStr parses a hotkey string and returns the according Hotkey's.
// Mods are unambiguous (though we ignore here the Num and Caps Lock).
// Keycodes aren't, though. For instance, on some layouts, there are 2 keys with different keycodes
// which map to 'backslash' AKA '\'. And we want to handle them both.
// This is why we return an array of Hotkey's here. And an error if the string is "naughty".
func HotkeysFromStr(str string) (hotkeys []Hotkey, err error) {
	mods, keys, err := keybind.ParseString(X, str)
	if err != nil {
		return nil, err
	}

	if mods&IgnoreMods > 0 {
		return nil, errors.New("hotkeys can't rely on the status of the Caps Lock or the Num Lock (Mod2)")
	}

	for _, key := range keys {
		hotkeys = append(hotkeys, Hotkey{mods, key})
	}
	return
}

// ToStr converts the Hotkey into its string representation.
// This helps us show in a human-firendly way the bound hotkeys in the 'info' command.
func (h Hotkey) ToStr() (str string) {
	str = keybind.ModifierString(h.mods)
	if str != "" {
		str += "-"
	}

	keysym := keybind.KeysymGet(X, h.key, 0)
	str += keybind.KeysymToBaseStr(keysym)
	return
}

func (h Hotkey) Grab() error {
	return keybind.GrabChecked(X, X.RootWin(), h.mods, h.key)
}

func (h Hotkey) Ungrab() {
	keybind.Ungrab(X, X.RootWin(), h.mods, h.key)
}

func HandleEvent(event xgb.Event) {
	switch ev := event.(type) {
	case xproto.KeyPressEvent:
		hotkey := Hotkey{
			mods: ev.State &^ IgnoreMods,
			key:  ev.Detail,
		}
		Exec(KeyPressCommands.Get(hotkey))
	case xproto.KeyReleaseEvent:
		hotkey := Hotkey{
			mods: ev.State &^ IgnoreMods,
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
