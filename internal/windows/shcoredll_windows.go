package windows

import (
	"syscall"
	"unsafe"

	"github.com/jholder85638/toolbox/errs"
)

var (
	shcore           = syscall.NewLazyDLL("shcore.dll")
	getDpiForMonitor = shcore.NewProc("GetDpiForMonitor")
)

// GetDpiForMonitor https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetDpiForMonitor(monitor HMONITOR, dpiType int32, dpiX, dpiY *uint32) error {
	ret, _, err := getDpiForMonitor.Call(uintptr(monitor), uintptr(dpiType), uintptr(unsafe.Pointer(dpiX)), uintptr(unsafe.Pointer(dpiY)))
	if ret != 0 {
		return errs.NewWithCause(getDpiForMonitor.Name, err)
	}
	return nil
}
