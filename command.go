// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import "context"

// Command implements a specific command.
type Command interface {
	Main(ctx context.Context, args []string) error
}

// CommandFunc transforms a func into a [Command].
type CommandFunc func(ctx context.Context, args []string) error

var _ Command = CommandFunc(nil)

// Main implements [Command].
func (fx CommandFunc) Main(ctx context.Context, args []string) error {
	return fx(ctx, args)
}
