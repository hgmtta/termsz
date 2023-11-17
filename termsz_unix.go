//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos

package termsz

import (
	"syscall"
	"unsafe"
)

func getSize() (cols, rows int, err error) {
	var sz struct {
		rows, cols, xpixels, ypixels uint16
	}
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&sz)),
	)
	if errno != 0 {
		err := syscall.Errno(errno)
		return 0, 0, err
	}
	return int(sz.cols), int(sz.rows), nil
}
