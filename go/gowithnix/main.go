package gowithnix

import (
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:     `gowithnix`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		fmt.Println("Go meets Nix")
		return nil
	},
}

var HelloCmd = &Z.Cmd{
	Name:     `hello`,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(caller *Z.Cmd, args ...string) error {
		fmt.Println("Hello, World!")
		return nil
	},
}
