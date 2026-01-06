// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"context"
	"errors"
	"fmt"

	"github.com/bassosimone/vflag"
)

// helpMain is the main of the help subcommand.
func (c *DispatcherCommand) helpMain(ctx context.Context, args []string) error {
	// initialize the flag set
	fset := vflag.NewFlagSet(fmt.Sprintf("%s help", c.Name), c.ErrorHandling)
	fset.UsagePrinter = c.NewHelpSubcommandUsagePrinter()
	fset.AutoHelp('h', "help", helpFlagDescr)
	fset.SetMinMaxPositionalArgs(0, 1)
	fset.Exit = c.Exit
	fset.Stderr = c.Stderr
	fset.Stdout = c.Stdout

	// parse the CLI arguments
	if err := fset.Parse(args); err != nil {
		if errors.Is(err, vflag.ErrHelp) {
			return c.printHelp()
		}
		fset.PrintUsageError(c.Stderr, err)
		return err
	}

	// check whether the user is requesting help for a subcommand
	if len(fset.Args()) > 0 {
		if cmd, ok := c.findCommand(fset.Args()[0]); ok {
			return cmd.Main(ctx, []string{"--help"})
		}
		err := fmt.Errorf("%w: %s", ErrCommandNotFound, args[0])
		fset.PrintUsageError(c.Stderr, err)
		return err
	}

	// print the general overall help
	return c.printHelp()
}

func (c *DispatcherCommand) printHelp() error {
	c.UsagePrinter.PrintHelp(c, c.Stdout)
	return nil
}
