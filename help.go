// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/bassosimone/must"
	"github.com/bassosimone/textwrap"
	"github.com/bassosimone/vflag"
)

// helpMain is the main of the help subcommand.
func (c *DispatcherCommand) helpMain(ctx context.Context, args []string) error {
	// initialize the flag set
	fset := vflag.NewFlagSet(fmt.Sprintf("%s help", c.Name), c.ErrorHandling)
	usage := vflag.NewDefaultUsagePrinter()
	usage.AddDescription(helpSubcommandDescr)
	fset.UsagePrinter = usage
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
	// print the help message
	const wrapAtColumn = 72

	// Usage
	must.Fprintf(c.Stdout, "\n")
	must.Fprintf(c.Stdout, "Usage\n")
	must.Fprintf(c.Stdout, "\n")
	for _, name := range slices.Sorted(maps.Keys(c.Commands)) {
		for _, alias := range c.CommandNameToAliases[name] {
			must.Fprintf(c.Stdout, "    %s %s [args...]\n", c.Name, alias)
		}
		must.Fprintf(c.Stdout, "    %s %s [args...]\n", c.Name, name)
		must.Fprintf(c.Stdout, "\n")
	}

	// Description
	if len(c.Description) > 0 {
		must.Fprintf(c.Stdout, "Description\n")
		for _, paragraph := range c.Description {
			must.Fprintf(c.Stdout, "\n")
			must.Fprintf(c.Stdout, "%s", textwrap.Do(paragraph, wrapAtColumn, "    "))
			must.Fprintf(c.Stdout, "\n")
		}
		must.Fprintf(c.Stdout, "\n")
	}

	// Commands
	must.Fprintf(c.Stdout, "Commands\n")
	for _, name := range slices.Sorted(maps.Keys(c.Commands)) {
		must.Fprintf(c.Stdout, "\n")
		aliases := slices.Clone(c.CommandNameToAliases[name])
		aliases = append(aliases, name)
		must.Fprintf(c.Stdout, "    %s\n", strings.Join(aliases, ", "))
		command := c.Commands[name]
		for _, paragraph := range command.descr {
			must.Fprintf(c.Stdout, "\n")
			must.Fprintf(c.Stdout, "%s", textwrap.Do(paragraph, wrapAtColumn, "        "))
			must.Fprintf(c.Stdout, "\n")
		}
	}

	// Hints
	must.Fprintf(c.Stdout, "\n")
	must.Fprintf(c.Stdout, "Hints\n")
	paragraphs := []string{
		fmt.Sprintf("Use `%s <command> --help' to get command-specific help.", c.Name),
		"Append `--help' or `-h' to any command line failing with usage errors to hide the " +
			"error and obtain contextual help.",
	}
	for _, paragraph := range paragraphs {
		must.Fprintf(c.Stdout, "\n")
		must.Fprintf(c.Stdout, "%s", textwrap.Do(paragraph, wrapAtColumn, "    "))
		must.Fprintf(c.Stdout, "\n")

	}

	must.Fprintf(c.Stdout, "\n")
	return nil
}
