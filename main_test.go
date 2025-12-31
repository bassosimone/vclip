// SPDX-License-Identifier: GPL-3.0-or-later

package vclip_test

import (
	"context"
	"testing"

	"github.com/bassosimone/vclip"
)

// We cannot easily test that [log.Fatal] is eventually
// called so, we settle for a simple smoke test.
func TestVclipMain(t *testing.T) {
	ctx := context.Background()
	command := vclip.CommandFunc(func(ctx context.Context, args []string) error {
		return nil
	})
	vclip.Main(ctx, command, []string{})
}
