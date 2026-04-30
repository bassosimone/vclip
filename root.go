//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/root.go
//

package vclip

import (
	"context"
	"os/signal"

	"github.com/bassosimone/runtimex"
)

// RootCommand implements the root command of a binary.
//
// Construct using [NewRootCommand].
type RootCommand struct {
	// Command is the [Command] to run.
	//
	// Set to the parameter passed to [NewRootCommand].
	Command Command

	// LogFatalOnError0 is the function to call when Command fails.
	//
	// [NewRootCommand] initializes it to [runtimex.LogFatalOnError0].
	LogFatalOnError0 func(err error)
}

// NewRootCommand creates and returns a new [*RootCommand] instance.
func NewRootCommand(cmd Command) *RootCommand {
	return &RootCommand{
		Command:          cmd,
		LogFatalOnError0: runtimex.LogFatalOnError0,
	}
}

// Main runs the underlying [Command] as follows:
//
// 1. we wrap the [context.Context] using [signal.NotifyContext] so that
// interruptions such as `^C` interrupt the command execution.
//
// 2. we use the LogFatalOnError0 field to ensure that the error
// returned by the command is logged and leads to exiting.
//
// The args MUST NOT contain the program name (i.e. use `os.Args[1:]`).
func (cmd *RootCommand) Main(ctx context.Context, args []string) {
	ctx, cancel := signal.NotifyContext(ctx, interruptSignals...)
	defer cancel()
	cmd.LogFatalOnError0(cmd.Command.Main(ctx, args))
}
