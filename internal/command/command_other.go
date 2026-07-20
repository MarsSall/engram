//go:build !windows

package command

import (
	"context"
	"os/exec"
)

// NewContext creates a command using the platform defaults.
func NewContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, args...)
}
