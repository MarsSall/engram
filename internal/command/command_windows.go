//go:build windows

package command

import (
	"context"
	"os/exec"
	"syscall"
)

// NewContext creates a command without showing a console window.
func NewContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
