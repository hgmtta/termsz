//go:build windows
// +build windows

package termsz

import (
	"syscall"
	"unsafe"
)

type (
	smallRect struct {
		Left, Top, Right, Bottom int16
	}
	coord struct {
		X, Y int16
	}
	bufferInfo struct {
		Size              coord
		CursorPosition    coord
		Attributes        uint16
		Window            smallRect
		MaximumWindowSize coord
	}
)

var kernel32DLL = syscall.NewLazyDLL("kernel32.dll")
var getBufferInfoProc = kernel32DLL.NewProc("GetConsoleScreenBufferInfo")

func getSize() (cols, rows int, err error) {
	stdoutHandle, err := getStdHandle(syscall.STD_OUTPUT_HANDLE)
	if err != nil {
		return 0, 0, err
	}

	var i *bufferInfo
	i, err = getBufferInfo(stdoutHandle)
	if err != nil {
		return 0, 0, err
	}

	return int(i.Window.Right - i.Window.Left + 1), int(i.Window.Bottom - i.Window.Top + 1), nil
}

func getStdHandle(stdhandle int) (uintptr, error) {
	handle, err := syscall.GetStdHandle(stdhandle)
	if err != nil {
		return 0, err
	}
	return uintptr(handle), nil
}

func getBufferInfo(handle uintptr) (*bufferInfo, error) {
	var i bufferInfo
	r1, _, err := getBufferInfoProc.Call(handle, uintptr(unsafe.Pointer(&i)), 0)
	if r1 == 0 {
		if err != nil {
			return nil, err
		}
		return nil, syscall.EINVAL
	}
	return &i, nil
}
