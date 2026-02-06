// SPDX-License-Identifier: GPL-3.0-or-later

package vclip_test

import (
	"context"
	"os"

	"github.com/bassosimone/vclip"
	"github.com/bassosimone/vflag"
)

// This example shows the usage printed when using the `--help` flag.
func Example_dispatcherCommandUsageWithHelpFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")

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
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	disp.AddDescription("Dispatcher for network commands.")

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
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	disp.AddDescription("Dispatcher for network commands.")

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
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	// Hints
	//
	//     Use `example <command> --help' to get command-specific help.
	//
	//     Append `--help' or `-h' to any command line failing with usage
	//     errors to hide the error and obtain contextual help.
}

// This example shows the version printed using the version command.
func Example_dispatcherCommandVersionCommand() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")
	disp.AddVersionHandlers("v0.1.0")

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

	// Invoke with `version` so that we print the version number
	disp.Main(ctx, []string{"version"})

	// Output:
	// v0.1.0
}

// This example shows the version printed using the `--version` flag.
func Example_dispatcherCommandVersionFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")
	disp.AddVersionHandlers("v0.1.0")

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

	// Invoke with `--version` so that we print the version number
	disp.Main(ctx, []string{"--version"})

	// Output:
	// v0.1.0
}

// This example shows that `version --help` just prints the version because
// the version subcommand ignores all arguments.
func Example_dispatcherCommandVersionHelp() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")
	disp.AddVersionHandlers("v0.1.0")

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

	// Invoke with `version --help`; the version subcommand ignores extra args
	disp.Main(ctx, []string{"version", "--help"})

	// Output:
	// v0.1.0
}

// This example shows that `help version` dispatches to `version --help`,
// which just prints the version since the version subcommand ignores args.
func Example_dispatcherCommandHelpVersion() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")
	disp.AddVersionHandlers("v0.1.0")

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

	// Invoke with `help version`; dispatches to `version --help` which prints the version
	disp.Main(ctx, []string{"help", "version"})

	// Output:
	// v0.1.0
}

// This example shows that `version --nope` just prints the version because
// the version subcommand ignores all arguments.
func Example_dispatcherCommandVersionInvalidFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")
	disp.AddVersionHandlers("v0.1.0")

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

	// Invoke with `version --nope`; the version subcommand ignores extra args
	disp.Main(ctx, []string{"version", "--nope"})

	// Output:
	// v0.1.0
}

// This example shows the error emitted for an invalid command
func Example_dispatcherCommandUsageWithInvalidCommand() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// Override Exit to transform it into a panic: using ExitOnError for a subcommand
	// eventually causes `disp.Exit(0)` to be invoked after printing help
	disp.Exit = func(status int) {
		panic("mocked exit invocation")
	}

	// Handle the panic by caused by Exit by simply ignoring it
	defer func() { recover() }()

	// Redirect the stderr to the stdout so that we can capture it
	disp.Stderr = os.Stdout

	// a background context is sufficient for this example
	ctx := context.Background()

	// Invoke with `nope` so that we fail
	disp.Main(ctx, []string{"nope"})

	// Output:
	// example: command not found: nope
	// example: use `example --help' to see the available commands
}

// This example shows that we show the help w/o failure w/ `--help` at the end of the command line.
func Example_dispatcherCommandUsageWithInvalidCommandIfAppendHelp() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// append --help at the end to neutralize the error and show the help
	disp.Main(ctx, []string{"nope", "--help"})

	// Output:
	// Usage
	//
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	// Hints
	//
	//     Use `example <command> --help' to get command-specific help.
	//
	//     Append `--help' or `-h' to any command line failing with usage
	//     errors to hide the error and obtain contextual help.
}

// This example shows that we show the help w/o failure w/ `-h` at the end of the command line.
func Example_dispatcherCommandUsageWithInvalidCommandIfAppendH() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// append -h at the end to neutralize the error and show the help
	disp.Main(ctx, []string{"nope", "-h"})

	// Output:
	// Usage
	//
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	// Hints
	//
	//     Use `example <command> --help' to get command-specific help.
	//
	//     Append `--help' or `-h' to any command line failing with usage
	//     errors to hide the error and obtain contextual help.
}

// This example shows what the help subcommand does when passed an invalid flag
// when using ContinueOnError (with ExitOnError we don't control what happens
// since it all depends on the vflag library)
func Example_dispatcherCommandUsageHelpWithInvalidFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// Redirect the stderr to the stdout so that we can capture it
	disp.Stderr = os.Stdout

	// a background context is sufficient for this example
	ctx := context.Background()

	// pass an invalid flag to `help` to see an error
	disp.Main(ctx, []string{"help", "--nope"})

	// Output:
	// example help: unknown option: --nope
	// example help: try `example help --help' for more help.
}

// This example shows that `help --help` is equivalent to `help` (if using ContinueOnError)
func Example_dispatcherCommandUsageHelpHelp() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// add an extra `--help` but we should see the same output
	disp.Main(ctx, []string{"help", "--help"})

	// Output:
	// Usage
	//
	//     example <command> [args...]
	//
	// Description
	//
	//     Dispatcher for network commands.
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
	// Hints
	//
	//     Use `example <command> --help' to get command-specific help.
	//
	//     Append `--help' or `-h' to any command line failing with usage
	//     errors to hide the error and obtain contextual help.
}

// This example shows what the help subcommand does when passed an invalid subcommand
// when using ContinueOnError (with ExitOnError we don't control what happens
// since it all depends on the vflag library)
func Example_dispatcherCommandUsageHelpWithInvalidSubcommand() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispatcher for network commands.")

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

	// Redirect the stderr to the stdout so that we can capture it
	disp.Stderr = os.Stdout

	// a background context is sufficient for this example
	ctx := context.Background()

	// pass an invalid subcommand to `help` to see an error
	disp.Main(ctx, []string{"help", "nope"})

	// Output:
	// example help: command not found: nope
	// example help: try `example help --help' for more help.
}
