package common

import "github.com/alexflint/go-arg"

type BindCmd struct {
	Released    bool   `arg:"-r" help:"bind the HOTKEY on its release instead of its press"`
	Hotkey      string `arg:"positional,required" help:"the hotkey to trigger the RUNCOMMAND"`
	RunCommand  string `arg:"positional,required" help:"the command to be run by 'sh' when the HOTKEY gets triggered"`
	Description string `arg:"positional" help:"an optional description to be shown by the 'info' subcommand"`
}

type UnbindCmd struct {
	Released bool   `arg:"-r" help:"unbind the HOTKEY's on-release command instead of the on-press one"`
	Hotkey   string `arg:"positional,required" help:"the hotkey to be unbound"`
}

type InfoCmd struct {
	Hotkey string `arg:"positional" help:"if set, show only info about the given HOTKEY"`
}

type Args struct {
	SocketPath       string   `arg:"-s,--socket" help:"use the given SOCKET file" default:"/tmp/gxhk.sock" json:"-"`
	ConfigFiles      []string `arg:"-c,--config,separate" help:"use the given CONFIG file, besides all the others" json:"-"`
	NoDefaultConfigs bool     `arg:"-C,--no-default-configs" help:"ignore the default config files" json:"-"`

	Bind   *BindCmd   `arg:"subcommand:bind" help:"bind keys to commands" json:",omitempty"`
	Unbind *UnbindCmd `arg:"subcommand:unbind" help:"unbind keys" json:",omitempty"`
	Info   *InfoCmd   `arg:"subcommand:info" help:"show infos about the bound keys and their commands" json:",omitempty"`
}

func ParseArgs() (Args, *arg.Parser) {
	var args Args
	p := arg.MustParse(&args)
	return args, p
}
