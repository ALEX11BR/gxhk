# gxhk

`gxhk` is a hotkey daemon for X, something like [`sxhkd`](https://github.com/baskerville/sxhkd). What makes it special is the fact that it can be configured on-the-fly the same way as [`bspwm`](https://github.com/baskerville/bspwm): through a socket in which commands to bind keys, to unbind previously bound keys, or even show infos about the bound keys can be sent.

## Usage
Hotkeys are case-insensitive, with this syntax: `modifier1-modifier2-...-key`. For more info on this, check [this link](https://pkg.go.dev/github.com/alex11br/xgbutil@v0.0.0-20211225011412-f2944427ac98/keybind#hdr-Key_sequence_format) of the underlaying library. Note that, for this moment, controling for the lock (Caps Lock) and the mod2 (Num Lock) modifiers isn't allowed.

The app uses by default some config files: `/etc/gxhkrc` and `$XDG_CONFIG_HOME/gxhk/gxhkrc`. Using the `-C` flag, they can be ignored, and with the `-c` flag, new ones can be specified. Using the `-s` flag, a different socket besides the default `/tmp/gxhk.sock` can be specified.

Here are some examples:
```sh
$ gxhk bind mod4-d 'rofi -show drun' # Spawn the rofi app menu when Super+d gets pressed
$ gxhk bind mod4-d 'rofi -show drun' 'Launch app menu' # The same thing, but with a description
$ gxhk bind Print 'maim ~/screenshot-$(date +%s).png && notify-send "Screenshot saved!"' # The command will be run using `sh -c`, so feel free to add variables, command chains, etc.
$ gxhk bind -r Mod4-Shift-Return '$TERMINAL' 'Launch terminal' # The terminal will be spawned when the hotkey Super+Shift+Enter gets released, not when it gets pressed as it happens with the ones above
$ gxhk info # Shows all the bound hotkeys so far
On mod4-d press: Launch app menu
On print press: Spawn 'maim ~/screenshot-$(date +%s).png && notify-send "Screenshot saved!"'
On mod4-shift-return release: Launch terminal
$ gxhk unbind Mod4-d # The Super+d hotkey will be unbound
$ gxhk info # Let's see what happens now
On print press: Spawn 'maim ~/screenshot-$(date +%s).png && notify-send "Screenshot saved!"'
On mod4-shift-return release: Launch terminal
```
You can check the [`gxhkrc`](gxhkrc) file for more inspiration, or, once you've installed the app, just copy the default config file like this: `install -Dm755 /etc/gxhkrc ~/.config/gxhk/gxhkrc` and start editing.
## Installation
### From source
Run `make build` to build the app, then `sudo make install` to install it.