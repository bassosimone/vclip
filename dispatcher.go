// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/bassosimone/must"
	"github.com/bassosimone/runtimex"
	"github.com/bassosimone/vflag"
)

// DispatcherCommand is a command that dispatches execution to subcommands.
//
// Construct using [NewDispatcherCommand].
type DispatcherCommand struct {
	// CommandAliasToName maps a command alias to its real name.
	//
	// [NewDispatcherCommand] initializes it as an empty map.
	CommandAliasToName map[string]string

	// CommandNameToAliases maps a command name to its aliases.
	//
	// [NewDispatcherCommand] initializes it as an empty map.
	CommandNameToAliases map[string][]string

	// Commands maps between names and [Command] instances.
	//
	// [NewDispatcherCommand] initializes the following built-in subcommands:
	//
	//  1. `-h`, `--help`, and `help` to print help.
	//
	//  2. `--version` and `version` to print the version number.
	Commands map[string]DescribedCommand

	// Description contains the description paragraphs.
	//
	// [NewDispatcherCommand] initializes this field to an empty slice.
	Description []string

	// ErrorHandling is the [vflag.ErrorHandling] policy to use.
	//
	// Set by the parameter passed to [NewDispatcherCommand].
	ErrorHandling vflag.ErrorHandling

	// Exit is the function to call with the [ExitOnError] policy.
	//
	// [NewDispatcherCommand] initializes it to [os.Exit].
	Exit func(status int)

	// Name is the command name.
	//
	// Set by the parameter passed to [NewDispatcherCommand].
	Name string

	// NewHelpSubcommandUsagePrinter returns the [vflag.UsagePrinter] that the
	// auto-generated help subcommand should use.
	//
	// [NewDispatcherCommand] initializes this function to a sane default
	// that includes a minimal command description.
	NewHelpSubcommandUsagePrinter func() vflag.UsagePrinter

	// NewVersionSubcommandUsagePrinter returns the [vflag.UsagePrinter] that the
	// auto-generated version subcommand should use.
	//
	// [NewDispatcherCommand] initializes this function to a sane default
	// that includes a minimal command description.
	NewVersionSubcommandUsagePrinter func() vflag.UsagePrinter

	// Stderr is the [io.Writer] to use as the stderr.
	//
	// [NewDispatcherCommand] initializes this field to [os.Stderr].
	//
	// We use this field with [ExitOnError] policy.
	Stderr io.Writer

	// Stdout is the [io.Writer] to use as the stdout.
	//
	// [NewDispatcherCommand] initializes this field to [os.Stdout].
	//
	// We use this field with [ExitOnError] policy.
	Stdout io.Writer

	// UsagePrinter is the [UsagePrinter] to use.
	//
	// Initialized by [NewDispatcherCommand] using [NewDefaultUsagePrinter].
	UsagePrinter UsagePrinter

	// Version contains the program version.
	//
	// [NewDispatcherCommand] initializes this field to "v0.0.0-dev".
	Version string
}

const (
	// versionSubcommandDescr describes the version subcommand and the --version flag.
	versionSubcommandDescr = "Show the version number and exit."

	// helpSubcommandDescr describes the --help and -h flags.
	helpFlagDescr = "Show this help message and exit."

	// helpSubcommandDescr describes the help subcommand and the --help and -h flags.
	helpSubcommandDescr = "Show help about this command or about a subcommand."
)

// NewDispatcherCommand creates a new instance of [*DispatcherCommand].
func NewDispatcherCommand(name string, handling vflag.ErrorHandling) *DispatcherCommand {
	c := &DispatcherCommand{
		CommandAliasToName:   map[string]string{},
		CommandNameToAliases: map[string][]string{},
		Commands:             map[string]DescribedCommand{},
		Description:          []string{},
		ErrorHandling:        handling,
		Exit:                 os.Exit,
		Name:                 name,
		NewHelpSubcommandUsagePrinter: func() vflag.UsagePrinter {
			usage := vflag.NewDefaultUsagePrinter()
			usage.AddDescription(helpSubcommandDescr)
			return usage
		},
		NewVersionSubcommandUsagePrinter: func() vflag.UsagePrinter {
			usage := vflag.NewDefaultUsagePrinter()
			usage.AddDescription(versionSubcommandDescr)
			return usage
		},
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		UsagePrinter: NewDefaultUsagePrinter(),
		Version:      "v0.0.0-dev",
	}

	c.AddCommand("help", CommandFunc(c.helpMain), helpSubcommandDescr)
	c.MustAddCommandAlias("help", "-h")
	c.MustAddCommandAlias("help", "--help")

	c.AddCommand("version", CommandFunc(c.versionMain), versionSubcommandDescr)
	c.MustAddCommandAlias("version", "--version")

	return c
}

// AddDescription adds the given text to the description paragraphs.
func (c *DispatcherCommand) AddDescription(text ...string) {
	c.Description = append(c.Description, text...)
}

// AddCommand adds a [Command] with the given name to the [*DispatcherCommand].
//
// Commands MUST handle the `--help` flag and provide help when they see it regardless
// of the otherwise different convention they use for flags.
func (c *DispatcherCommand) AddCommand(name string, cmd Command, descr ...string) {
	c.Commands[name] = NewDescribedCommand(cmd, descr...)
}

// MustAddCommandAlias introduces an alias for an existing command.
//
// This method panics if curName is not an existing command name.
func (c *DispatcherCommand) MustAddCommandAlias(curName, newAlias string) {
	_, found := c.Commands[curName]
	runtimex.Assert(found)
	c.CommandAliasToName[newAlias] = curName
	c.CommandNameToAliases[curName] = append(c.CommandNameToAliases[curName], newAlias)
}

// findCommand searches for a command taking aliases into account.
func (c *DispatcherCommand) findCommand(name string) (Command, bool) {
	if realName, ok := c.CommandAliasToName[name]; ok {
		name = realName
	}
	cmd, ok := c.Commands[name]
	return cmd, ok
}

var _ Command = &DispatcherCommand{}

// ErrCommandNotFound indicates that the given command was not found.
var ErrCommandNotFound = errors.New("command not found")

// Main implements [CommandHandler].
func (c *DispatcherCommand) Main(ctx context.Context, args []string) error {
	return c.maybeHandleError(c.main(ctx, args))
}

func (c *DispatcherCommand) main(ctx context.Context, args []string) error {
	if len(args) <= 0 {
		return c.helpMain(ctx, args)
	}
	if child, ok := c.findCommand(args[0]); ok {
		return child.Main(ctx, args[1:])
	}
	return c.maybeRecoverErrCommandNotFound(args)
}

func (c *DispatcherCommand) maybeHandleError(err error) error {
	// Determine what to do based on the policy
	switch {
	case err == nil:
		return nil

	case c.ErrorHandling == vflag.ContinueOnError:
		return err

	case c.ErrorHandling == vflag.ExitOnError:
		must.Fprintf(c.Stderr, "%s\n", err.Error())
		switch {
		case errors.Is(err, ErrCommandNotFound):
			must.Fprintf(c.Stderr, "%s: use `%s --help' to see the available commands\n", c.Name, c.Name)
			c.Exit(2)
		default:
			c.Exit(1)
		}
	}

	// We end up here for [PanicOnError] or whenever c.Exit is so
	// broken that it doesn't actually exit.
	panic(err)
}

// maybeRecoverErrCommandNotFound recovers from an [ErrCommandNotFound] error condition
// when the `--help` or the `-h` flag appears at the end of the command line.
//
// This specifically enables the UX pattern where the user appends `-h` or `--help` to
// the command line, however wrong, and always gets the usage.
func (c *DispatcherCommand) maybeRecoverErrCommandNotFound(args []string) error {
	total := len(args)
	if total < 1 || (args[total-1] != "--help" && args[total-1] != "-h") {
		return fmt.Errorf("%s: %w: %s", c.Name, ErrCommandNotFound, args[0])
	}
	return c.printHelp()
}
