// SPDX-License-Identifier: GPL-3.0-or-later

package vclip_test

import (
	"context"

	"github.com/bassosimone/vclip"
	"github.com/bassosimone/vflag"
)

// This example shows the usage printed when using the `--help` flag.
func Example_dispatcherCommandUsageWithHelpFlag() {
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

// This example shows the usage printed when invoked without arguments.
func Example_dispatcherCommandUsageWithoutArguments() {
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

	// Invoke without argumuments so that we print the help
	disp.Main(ctx, []string{})

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

// This example shows the usage printed when using the `-h` flag.
func Example_dispatcherCommandUsageWithHFlag() {
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

	// Invoke with `-h` so that we print the help
	disp.Main(ctx, []string{"-h"})

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
