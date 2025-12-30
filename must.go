// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"fmt"
	"io"

	"github.com/bassosimone/runtimex"
)

// mustFprintf is like [fmt.Fprintf] but panics in case of failure.
func mustFprintf(w io.Writer, format string, args ...any) {
	// TODO(bassosimone): this should probably live in its own tiny library
	_ = runtimex.PanicOnError1(fmt.Fprintf(w, format, args...))
}
