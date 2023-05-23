package main

import (
	"flag"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

type Command struct {
	Z     *Z.Cmd
	Flags *flag.FlagSet
}

var Cmd = &Command{
	Z: &Z.Cmd{
		Name:     `goflag`,
		Commands: []*Z.Cmd{help.Cmd},
		Call: func(caller *Z.Cmd, args ...string) error {
			return nil
		},
	},
}

var name string

func main() {
	Cmd.Flags = flag.NewFlagSet(Cmd.Z.Name, flag.ExitOnError)
	Cmd.Flags.StringVar(&name, "name", "", "name of invoker")
}
