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

// This example shows the version printed using the version command.
func Example_dispatcherCommandVersionCommand() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddDescription("Dispacher for network commands.")
	disp.Version = "v0.1.0"

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
	disp.AddDescription("Dispacher for network commands.")
	disp.Version = "v0.1.0"

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

// This example shows the usage printed using the `version --help` arguments
// when using ContinueOnError (ExitOnError exits immediately so it's a different
// codepath not controller by this library.)
func Example_dispatcherCommandVersionHelp() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispacher for network commands.")
	disp.Version = "v0.1.0"

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

	// Invoke with `version --help` so that we print the version subcommand usage
	disp.Main(ctx, []string{"version", "--help"})

	// Output:
	// Usage
	//
	//     example version [flags]
	//
	// Description
	//
	//     Show the version number and exit.
	//
	// Flags
	//
	//     -h, --help
	//
	//         Show this help message and exit.
}

// This example shows that `help version` is equivalent to `version --help`
// when using ContinueOnError (ExitOnError exits immediately so it's a different
// codepath not controller by this library.)
func Example_dispatcherCommandHelpVersion() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispacher for network commands.")
	disp.Version = "v0.1.0"

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

	// Invoke with `help version` so that we print the version subcommand usage
	disp.Main(ctx, []string{"help", "version"})

	// Output:
	// Usage
	//
	//     example version [flags]
	//
	// Description
	//
	//     Show the version number and exit.
	//
	// Flags
	//
	//     -h, --help
	//
	//         Show this help message and exit.
}

// This example shows an error being printed with invalid flag passed to `version`
// when using ContinueOnError (ExitOnError exits immediately so it's a different
// codepath not controller by this library.)
func Example_dispatcherCommandVersionInvalidFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddDescription("Dispacher for network commands.")
	disp.Version = "v0.1.0"

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

	// Invoke with `help version` so that we print the version subcommand usage
	disp.Main(ctx, []string{"version", "--nope"})

	// Output:
	// example version: unknown option: --nope
	// example version: try `example version --help' for more help.
}

// This example shows the error emitted for an invalid command
func Example_dispatcherCommandUsageWithInvalidCommand() {
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

	// append --help at the end to neutralize the error and show the help
	disp.Main(ctx, []string{"nope", "--help"})

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

// This example shows that we show the help w/o failure w/ `-h` at the end of the command line.
func Example_dispatcherCommandUsageWithInvalidCommandIfAppendH() {
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

	// append -h at the end to neutralize the error and show the help
	disp.Main(ctx, []string{"nope", "-h"})

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

// This example shows what the help subcommand does when passed an invalid flag
// when using ContinueOnError (with ExitOnError we don't control what happens
// since it all depends on the vflag library)
func Example_dispatcherCommandUsageHelpWithInvalidFlag() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
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

	// add an extra `--help` but we should see the same output
	disp.Main(ctx, []string{"help", "--help"})

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

// This example shows what the help subcommand does when passed an invalid subcommand
// when using ContinueOnError (with ExitOnError we don't control what happens
// since it all depends on the vflag library)
func Example_dispatcherCommandUsageHelpWithInvalidSubcommand() {
	// create and init the dispatcher command
	disp := vclip.NewDispatcherCommand("example", vflag.ContinueOnError)
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
