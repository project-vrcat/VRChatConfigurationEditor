package win32

import (
	"syscall"
	"unsafe"
)

var (
	shell32                 = syscall.NewLazyDLL("shell32.dll")
	procSHBrowseForFolder   = shell32.NewProc("SHBrowseForFolderW")
	procSHGetPathFromIDList = shell32.NewProc("SHGetPathFromIDListW")
)

// SHBrowseForFolder flags
const (
	BIF_RETURNONLYFSDIRS    = 0x00000001
	BIF_DONTGOBELOWDOMAIN   = 0x00000002
	BIF_STATUSTEXT          = 0x00000004
	BIF_RETURNFSANCESTORS   = 0x00000008
	BIF_EDITBOX             = 0x00000010
	BIF_VALIDATE            = 0x00000020
	BIF_NEWDIALOGSTYLE      = 0x00000040
	BIF_BROWSEINCLUDEURLS   = 0x00000080
	BIF_USENEWUI            = BIF_EDITBOX | BIF_NEWDIALOGSTYLE
	BIF_UAHINT              = 0x00000100
	BIF_NONEWFOLDERBUTTON   = 0x00000200
	BIF_NOTRANSLATETARGETS  = 0x00000400
	BIF_BROWSEFORCOMPUTER   = 0x00001000
	BIF_BROWSEFORPRINTER    = 0x00002000
	BIF_BROWSEINCLUDEFILES  = 0x00004000
	BIF_SHAREABLE           = 0x00008000
	BIF_BROWSEFILEJUNCTIONS = 0x00010000
)

func SHGetPathFromIDList(idl uintptr) string {
	buf := make([]uint16, 1024)
	_, _, _ = procSHGetPathFromIDList.Call(
		idl,
		uintptr(unsafe.Pointer(&buf[0])),
	)
	return syscall.UTF16ToString(buf)
}

func SHBrowseForFolder(bi *BROWSEINFO) uintptr {
	ret, _, _ := procSHBrowseForFolder.Call(uintptr(unsafe.Pointer(bi)))
	return ret
}
