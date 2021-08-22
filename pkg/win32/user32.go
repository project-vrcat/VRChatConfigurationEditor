package win32

import (
	"syscall"
	"unsafe"
)

var (
	user32                       = syscall.NewLazyDLL("user32.dll")
	procEnumWindows              = user32.NewProc("EnumWindows")
	procMessageBox               = user32.NewProc("MessageBoxW")
	procShowWindow               = user32.NewProc("ShowWindow")
	procIsWindowVisible          = user32.NewProc("IsWindowVisible")
	procGetWindowThreadProcessId = user32.NewProc("GetWindowThreadProcessId")
)

// MessageBox flags
const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300
	MB_TOPMOST           = 0x00040000
)

const (
	IDOK       = 1
	IDCANCEL   = 2
	IDABORT    = 3
	IDRETRY    = 4
	IDIGNORE   = 5
	IDYES      = 6
	IDNO       = 7
	IDCLOSE    = 8
	IDHELP     = 9
	IDTRYAGAIN = 10
	IDCONTINUE = 11
	IDTIMEOUT  = 32000
)

func MessageBox(hwnd HWND, text, caption string, flags uint) int {
	lpText, _ := syscall.UTF16PtrFromString(text)
	lpCaption, _ := syscall.UTF16PtrFromString(caption)
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		uintptr(flags),
	)
	return int(ret)
}

func ShowWindow(hwnd HWND, cmdShow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdShow),
	)
	return ret != 0
}

func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
	var processId DWORD
	ret, _, _ := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)),
	)
	return HANDLE(ret), int(processId)
}

func EnumWindows(enumFunc uintptr, lParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumWindows.Addr(), 2, enumFunc, lParam, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func IsWindowVisible(hwnd HWND) bool {
	ret, _, _ := procIsWindowVisible.Call(uintptr(hwnd))
	return ret != 0
}

func FindWindowByProcessId(pid int) HWND {
	_hwnd := HWND(0)
	cb := syscall.NewCallback(func(hwnd syscall.Handle, lParam uintptr) uintptr {
		_, _pid := GetWindowThreadProcessId(HWND(hwnd))
		if _pid == pid && IsWindowVisible(HWND(hwnd)) {
			_hwnd = HWND(hwnd)
			return 0
		}
		return 1
	})
	_ = EnumWindows(cb, 0)
	return _hwnd
}
