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

// Main runs the given [Command] wrapping its execution as follows:
//
// 1. we wrap the [context.Context] using [signal.NotifyContext] so that
// interruptions such as `^C` interrupt the command exectution.
//
// 2. we use [runtimex.LogFatalOnError0] to ensure that the error
// returned by the command is logged and leads to exiting.
//
// The args MUST NOT contain the program name (e.g. `os.Args[1:]`).
func Main(ctx context.Context, cmd Command, args []string) {
	ctx, cancel := signal.NotifyContext(ctx, interruptSignals...)
	defer cancel()
	runtimex.PanicOnError0(cmd.Main(ctx, args))
}
