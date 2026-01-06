// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"context"
	"errors"
	"fmt"

	"github.com/bassosimone/must"
	"github.com/bassosimone/vflag"
)

// versionMain is the main of the version subcommand.
func (c *DispatcherCommand) versionMain(ctx context.Context, args []string) error {
	// initialize the flag set
	fset := vflag.NewFlagSet(fmt.Sprintf("%s version", c.Name), c.ErrorHandling)
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription(versionSubcommandDescr)
	fset.UsagePrinter = usage
	fset.AutoHelp('h', "help", helpFlagDescr)
	fset.Exit = c.Exit
	fset.Stderr = c.Stderr
	fset.Stdout = c.Stdout

	// parse the CLI arguments
	if err := fset.Parse(args); err != nil {
		if errors.Is(err, vflag.ErrHelp) {
			fset.PrintUsageString(c.Stdout)
			return nil
		}
		fset.PrintUsageError(c.Stderr, err)
		return err
	}

	// print the version on stdout
	must.Fprintf(c.Stdout, "%s\n", c.Version)
	return nil
}
