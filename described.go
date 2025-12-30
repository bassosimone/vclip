// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import "context"

// describedCommand is a command living side by side with its docs.
type describedCommand struct {
	// cmd is the [Command].
	cmd Command

	// descr contains the [Command] description.
	descr []string
}

// newDescribedCommand creates a [Command] along with the related documentation.
func newDescribedCommand(cmd Command, descr ...string) describedCommand {
	return describedCommand{
		cmd:   cmd,
		descr: descr,
	}
}

var _ Command = describedCommand{}

// Main implements [Command].
func (dc describedCommand) Main(ctx context.Context, args []string) error {
	return dc.cmd.Main(ctx, args)
}
