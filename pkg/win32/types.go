package win32

type (
	DWORD  uint32
	HANDLE uintptr
	HWND   HANDLE
)

// BROWSEINFO http://msdn.microsoft.com/en-us/library/windows/desktop/bb773205.aspx
type BROWSEINFO struct {
	Owner        HWND
	Root         *uint16
	DisplayName  *uint16
	Title        *uint16
	Flags        uint32
	CallbackFunc uintptr
	LParam       uintptr
	Image        int32
}
