//go:build linux
// +build linux

// syscall unix.Reboot is missing on macOS
package main

import "golang.org/x/sys/unix"

func main() {
	unix.Sync()

	unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART)
}
