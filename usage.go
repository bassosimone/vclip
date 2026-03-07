// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"

	"github.com/bassosimone/must"
	"github.com/bassosimone/textwrap"
)

// UsagePrinter prints the help for [*DispatcherCommand].
type UsagePrinter interface {
	PrintHelp(c *DispatcherCommand, w io.Writer)
}

// Constants controlling text formatting.
const (
	wrapAtColumn = 72
	indent4      = "    "
	indent8      = indent4 + indent4
)

// DefaultUsagePrinter is the default [UsagePrinter] implementation.
//
// Construct using [NewDefaultUsagePrinter].
type DefaultUsagePrinter struct{}

// NewDefaultUsagePrinter constructs a new [*DefaultUsagePrinter].
func NewDefaultUsagePrinter() *DefaultUsagePrinter {
	return &DefaultUsagePrinter{}
}

var _ UsagePrinter = &DefaultUsagePrinter{}

// PrintHelp implements [UsagePrinter].
//
// This method panics on I/O error.
func (up *DefaultUsagePrinter) PrintHelp(c *DispatcherCommand, w io.Writer) {
	// ## Usage
	must.Fprintf(w, "\n")
	must.Fprintf(w, "Usage\n")
	must.Fprintf(w, "\n")
	must.Fprintf(w, "    %s <command> [args...]\n", c.Name)
	must.Fprintf(w, "\n")

	// ## Description
	if len(c.Description) > 0 {
		must.Fprintf(w, "Description\n")
		for _, paragraph := range c.Description {
			up.div1(w, paragraph)
		}
		must.Fprintf(w, "\n")
	}

	// ## Commands
	must.Fprintf(w, "Commands\n")
	for _, name := range slices.Sorted(maps.Keys(c.Commands)) {
		must.Fprintf(w, "\n")
		aliases := slices.Clone(c.CommandNameToAliases[name])
		aliases = append(aliases, name)
		must.Fprintf(w, "    %s\n", strings.Join(aliases, ", "))
		command := c.Commands[name]
		for _, paragraph := range command.Descr {
			up.div2(w, paragraph)
		}
	}

	// ## Hints
	must.Fprintf(w, "\n")
	must.Fprintf(w, "Hints\n")
	paragraphs := []string{
		fmt.Sprintf("Use `%s <command> --help' to get command-specific help.", c.Name),
		"Append `--help' or `-h' to any command line failing with usage errors to hide the " +
			"error and obtain contextual help.",
	}
	for _, paragraph := range paragraphs {
		up.div1(w, paragraph)
	}

	must.Fprintf(w, "\n")
}

// div1 prints a paragraph at 4-space indent level. If the paragraph starts
// with 4 spaces, it is emitted verbatim (to allow preformatted blocks).
func (up *DefaultUsagePrinter) div1(w io.Writer, entry string) {
	must.Fprintf(w, "\n")
	if strings.HasPrefix(entry, indent4) {
		must.Fprintf(w, "%s%s\n", indent4, entry)
	} else {
		must.Fprintf(w, "%s\n", textwrap.Do(entry, wrapAtColumn, indent4))
	}
}

// div2 prints a paragraph at 8-space indent level. If the paragraph starts
// with 4 spaces, it is emitted verbatim (to allow preformatted blocks).
func (up *DefaultUsagePrinter) div2(w io.Writer, entry string) {
	must.Fprintf(w, "\n")
	if strings.HasPrefix(entry, indent4) {
		must.Fprintf(w, "%s%s\n", indent8, entry)
	} else {
		must.Fprintf(w, "%s\n", textwrap.Do(entry, wrapAtColumn, indent8))
	}
}
