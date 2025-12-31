// SPDX-License-Identifier: GPL-3.0-or-later

// Package vclip implements a minimal, versatile command-line dispatcher.
//
// The core idea is to keep the top-level parser small and neutral: it routes to
// subcommands but does not try to interpret or reorder their flags. This allows
// different subcommands to adopt different flag conventions while still sharing
// a consistent entry point and help behavior.
//
// Design tradeoffs and behavior:
//
//   - No global flag parsing: arguments before the subcommand are not reshuffled
//     or guessed, because ownership is ambiguous when subcommands have different
//     conventions.
//
//   - Help is universal: the dispatcher treats "-h" and "--help" as aliases for
//     the built-in help command, and appending "-h/--help" to a failing top-level
//     invocation turns the error into contextual help.
//
//   - Subcommands are responsible for their own flags: each command can use any
//     parsing style (e.g., vflag), as long as it honors "--help".
//
// The package evolved from real-world CLI use cases where subcommands emulate
// tools such as curl and dig, and it is designed to make migrating mixed-style
// parsers straightforward.
package vclip
