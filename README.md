# Golang Versatile CLI Command Dispatcher

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/vclip)](https://pkg.go.dev/github.com/bassosimone/vclip) [![Build Status](https://github.com/bassosimone/vclip/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/vclip/actions) [![codecov](https://codecov.io/gh/bassosimone/vclip/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/vclip)

The `vclip` Go package contains a minimal dispatcher for
command-line tools with subcommands. It keeps the top-level parser
neutral and lets each subcommand implement its own flag style.

For example:

```Go
import (
	"context"

	"github.com/bassosimone/runtimex"
	"github.com/bassosimone/vclip"
	"github.com/bassosimone/vflag"
)

// Create dispatcher
disp := vclip.NewDispatcherCommand("example", vflag.ExitOnError)
disp.AddDescription("Dispatcher for network commands.")

// Add curl subcommand
disp.AddCommand(
	"curl",
	vclip.CommandFunc(func(ctx context.Context, args []string) error {
		fset := vflag.NewFlagSet("example ls", vflag.ExitOnError)
		fLocation := false
		fset.BoolVar(&fLocation, 'l', "location")
		fSilent := false
		fset.BoolVar(&fSilent, 's', "silent")
		runtimex.PanicOnError0(fset.Parse(args))
		// ...
	}),
	"Utility to transfer URLs.",
)

// Add dig subcommand
disp.AddCommand(
	"dig",
	vclip.CommandFunc(func(ctx context.Context, args []string) error {
		// ...
	}),
	"Utility to query DNS servers.",
)

// Invoke with command line arguments (i.e. without the program name)
vclip.Main(context.Background(), disp, []string{"--help"})
```

The above example only sketches the setup; see [example_test.go](example_test.go)
for complete usage and output.

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/vclip
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Heavily refactored from original code at [bassosimone/clip](https://github.com/bassosimone/clip/tree/v0.8.0).
