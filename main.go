//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/root.go
//

package vclip

import "context"

// Main runs the given [Command] wrapping its execution as follows:
//
// 1. we wrap the [context.Context] using [signal.NotifyContext] so that
// interruptions such as `^C` interrupt the command execution.
//
// 2. we use [runtimex.LogFatalOnError0] to ensure that the error
// returned by the command is logged and leads to exiting.
//
// The args MUST NOT contain the program name (i.e. use `os.Args[1:]`).
//
// Use [*RootCommand] when you need to override the unconditional
// [runtimex.LogFatalOnError0] invocation executed by this function
// in case the underlying command fails.
func Main(ctx context.Context, cmd Command, args []string) {
	NewRootCommand(cmd).Main(ctx, args)
}
