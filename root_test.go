// SPDX-License-Identifier: GPL-3.0-or-later

package vclip_test

import (
	"context"
	"errors"
	"testing"

	"github.com/bassosimone/vclip"
	"github.com/stretchr/testify/require"
)

func TestRootCommandMain_Success(t *testing.T) {
	ctx := context.Background()
	command := vclip.CommandFunc(func(ctx context.Context, args []string) error {
		return nil
	})
	root := vclip.NewRootCommand(command)
	var gotErr error
	root.LogFatalOnError0 = func(err error) {
		gotErr = err
	}
	root.Main(ctx, []string{})
	require.Nil(t, gotErr)
}

func TestRootCommandMain_Failure(t *testing.T) {
	ctx := context.Background()
	expectedErr := errors.New("mocked error")
	command := vclip.CommandFunc(func(ctx context.Context, args []string) error {
		return expectedErr
	})
	root := vclip.NewRootCommand(command)
	var gotErr error
	root.LogFatalOnError0 = func(err error) {
		gotErr = err
	}
	root.Main(ctx, []string{})
	require.Equal(t, expectedErr, gotErr)
}
