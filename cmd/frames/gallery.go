package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arran4/golang-frame/cli"
)

var _ Cmd = (*galleryCmd)(nil)

type galleryCmd struct {
	*RootCmd
	Flags *flag.FlagSet

	SubCommands map[string]Cmd
}

func (c *galleryCmd) Usage() {
	err := executeUsage(os.Stderr, "gallery_usage.txt", c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating usage: %s\n", err)
	}
}

func (c *galleryCmd) Execute(args []string) error {
	if len(args) > 0 {
		if cmd, ok := c.SubCommands[args[0]]; ok {
			return cmd.Execute(args[1:])
		}
	}
	err := c.Flags.Parse(args)
	if err != nil {
		return NewUserError(err, fmt.Sprintf("flag parse error %s", err.Error()))
	}
	if err := cli.Gallery(); err != nil {
		return fmt.Errorf("gallery failed: %w", err)
	}
	return nil
}

func (c *RootCmd) NewgalleryCmd() *galleryCmd {
	set := flag.NewFlagSet("gallery", flag.ContinueOnError)
	v := &galleryCmd{
		RootCmd:     c,
		Flags:       set,
		SubCommands: make(map[string]Cmd),
	}

	set.Usage = v.Usage

	return v
}
