package main

import "github.com/alexflint/go-arg"

type BindCmd struct {
	Released    bool   `arg:"-r"`
	Hotkey      string `arg:"positional,required"`
	RunCommand  string `arg:"positional,required"`
	Description string `arg:"positional"`
}

type UnbindCmd struct {
	Released bool   `arg:"-r"`
	Hotkey   string `arg:"positional"`
}

type InfoCmd struct {
	Hotkey string `arg:"positional"`
}

type Args struct {
	SocketPath  string   `default:"/tmp/gxhk.sock" json:"-"`
	ConfigFiles []string `arg:"-c,separate" json:"-"`

	Bind   *BindCmd   `arg:"subcommand:bind" json:",omitempty"`
	Unbind *UnbindCmd `arg:"subcommand:unbind" json:",omitempty"`
	Info   *InfoCmd   `arg:"subcommand:info" json:",omitempty"`
}

func ParseArgs() (Args, *arg.Parser) {
	var args Args
	p := arg.MustParse(&args)
	return args, p
}
