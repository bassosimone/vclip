//go:build unix

//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/bassosimone/clip/blob/v0.8.0/signals_unix.go
//

package vclip

import (
	"os"
	"syscall"
)

var interruptSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
