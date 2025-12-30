// SPDX-License-Identifier: GPL-3.0-or-later

package vclip_test

import (
	"context"

	"github.com/bassosimone/vclip"
	"github.com/bassosimone/vflag"
)

// This example shows the usage printed by the dispatcher command.
func Example_dispatcherCommandUsage() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispacher for network commands.")

	// add two commands faking curl and dig
	disp.AddCommand(
		"curl",
		vclip.CommandFunc(func(ctx context.Context, args []string) error {
			return nil
		}),
		"Utility to transfer URLs.",
	)
	disp.AddCommand(
		"dig",
		vclip.CommandFunc(func(ctx context.Context, args []string) error {
			return nil
		}),
		"Utility to query DNS servers.",
	)

	// a background context is sufficient for this example
	ctx := context.Background()

	// override [os.Exit] to be panic to prevent the example from exiting
	disp.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Invoke with `--help` so that we print the help
	disp.Main(ctx, []string{"--help"})

	// Output:
	// Usage
	//
	//     example curl [args...]
	//
	//     example dig [args...]
	//
	//     example -h [args...]
	//     example --help [args...]
	//     example help [args...]
	//
	//     example --version [args...]
	//     example version [args...]
	//
	// Description
	//
	//     Dispacher for network commands.
	//
	// Commands
	//
	//     curl
	//
	//         Utility to transfer URLs.
	//
	//     dig
	//
	//         Utility to query DNS servers.
	//
	//     -h, --help, help
	//
	//         Show help about this command or about a subcommand.
	//
	//     --version, version
	//
	//         Show the version number and exit.
	//
	// Hints
	//
	//     Use `example <command> --help' to get command-specific help.
	//
	//     Append `--help' or `-h' to any command line failing with usage
	//     errors to hide the error and obtain contextual help.
}
