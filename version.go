// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"context"

	"github.com/bassosimone/must"
)

// versionMainFunc return the function to implement the `version` subcommand.
func (c *DispatcherCommand) versionMainFunc(version string) func(ctx context.Context, args []string) error {
	return func(ctx context.Context, args []string) error {
		must.Fprintf(c.Stdout, "%s\n", version)
		return nil
	}
}
