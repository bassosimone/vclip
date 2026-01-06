// SPDX-License-Identifier: GPL-3.0-or-later

package vclip

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/bassosimone/vflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCommand struct{}

func (tc *testCommand) Main(ctx context.Context, args []string) error {
	return nil
}

func TestDispatcherCommandFindCommandRoundTrip(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)
	cmd := &testCommand{}
	disp.AddCommand("echo", cmd, "echo text")
	disp.MustAddCommandAlias("echo", "e")

	found, ok := disp.findCommand("echo")
	require.True(t, ok)
	foundAlias, ok := disp.findCommand("e")
	require.True(t, ok)

	foundDescr, ok := found.(describedCommand)
	require.True(t, ok)
	foundAliasDescr, ok := foundAlias.(describedCommand)
	require.True(t, ok)

	assert.Same(t, cmd, foundDescr.cmd)
	assert.Same(t, cmd, foundAliasDescr.cmd)
}

type exitPanic struct {
	status int
}

func runMainExpectExit(t *testing.T, disp *DispatcherCommand, args []string) (int, string, string) {
	t.Helper()

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	disp.Stdout = &stdout
	disp.Stderr = &stderr
	disp.Exit = func(status int) { panic(exitPanic{status: status}) }

	var (
		status    int
		recovered bool
	)
	func() {
		defer func() {
			if r := recover(); r != nil {
				ep, ok := r.(exitPanic)
				require.True(t, ok)
				status = ep.status
				recovered = true
			}
		}()
		_ = disp.Main(context.Background(), args)
	}()

	require.True(t, recovered, "expected Exit to be invoked")
	return status, stdout.String(), stderr.String()
}

func TestDispatcherCommandMainContinueOnErrorReturnsError(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)
	sentinel := errors.New("boom")
	disp.AddCommand("fail", CommandFunc(func(ctx context.Context, args []string) error {
		return sentinel
	}))

	err := disp.Main(context.Background(), []string{"fail"})
	require.Error(t, err)
	assert.ErrorIs(t, err, sentinel)
}

func TestDispatcherCommandMainNoError(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)
	disp.AddCommand("ok", CommandFunc(func(ctx context.Context, args []string) error {
		return nil
	}))

	err := disp.Main(context.Background(), []string{"ok"})
	require.NoError(t, err)
}

func TestDispatcherCommandMainExitOnErrorCommandNotFound(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ExitOnError)
	status, _, stderr := runMainExpectExit(t, disp, []string{"nope"})

	assert.Equal(t, 2, status)
	assert.Contains(t, stderr, "example: use `example --help'")
}

func TestDispatcherCommandMainExitOnErrorCommandError(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ExitOnError)
	disp.AddCommand("fail", CommandFunc(func(ctx context.Context, args []string) error {
		return errors.New("boom")
	}))

	status, _, stderr := runMainExpectExit(t, disp, []string{"fail"})

	assert.Equal(t, 1, status)
	assert.Contains(t, stderr, "boom")
}

func TestDispatcherCommandMainPanicOnError(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.PanicOnError)
	sentinel := errors.New("boom")
	disp.AddCommand("fail", CommandFunc(func(ctx context.Context, args []string) error {
		return sentinel
	}))

	assert.PanicsWithValue(t, sentinel, func() {
		_ = disp.Main(context.Background(), []string{"fail"})
	})
}

func TestDispatcherCommandMainRecoverErrCommandNotFoundWithHelpFlag(t *testing.T) {
	testCases := []string{"--help", "-h"}
	for _, lastArg := range testCases {
		disp := NewDispatcherCommand("example", vflag.ContinueOnError)
		var stdout bytes.Buffer
		disp.Stdout = &stdout

		err := disp.Main(context.Background(), []string{"nope", lastArg})
		require.NoError(t, err)
		assert.Contains(t, stdout.String(), "Usage")
	}
}

func TestDispatcherCommandHelpMainUnknownFlag(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)

	err := disp.Main(context.Background(), []string{"help", "--bogus"})
	require.Error(t, err)
}

func TestDispatcherCommandHelpMainUnknownCommand(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)

	err := disp.Main(context.Background(), []string{"help", "nope"})
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrCommandNotFound)
}

func TestDispatcherCommandVersionMainUnknownFlag(t *testing.T) {
	disp := NewDispatcherCommand("example", vflag.ContinueOnError)

	err := disp.Main(context.Background(), []string{"version", "--bogus"})
	require.Error(t, err)
}
