// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import "context"

// DescribedCommand is a command living side by side with its docs.
type DescribedCommand struct {
	// Cmd is the [Command].
	Cmd Command

	// Descr contains the [Command] description.
	Descr []string
}

// NewDescribedCommand creates a [Command] along with the related documentation.
func NewDescribedCommand(cmd Command, descr ...string) DescribedCommand {
	return DescribedCommand{
		Cmd:   cmd,
		Descr: descr,
	}
}

var _ Command = DescribedCommand{}

// Main implements [Command].
func (dc DescribedCommand) Main(ctx context.Context, args []string) error {
	return dc.Cmd.Main(ctx, args)
}
