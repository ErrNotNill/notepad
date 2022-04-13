package windows

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unsafe"
)

//this block is responsible for Change Wallpaper on Windows
var (
	User32DLL           = windows.NewLazyDLL("user32.dll")
	procSystemParamInfo = User32DLL.NewProc("SystemParametersInfoW")
)

func SetWallpaper() {
	imagePath, _ := windows.UTF16PtrFromString(`C:\image.jpg`)
	fmt.Println("[+] Changing background now...")
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePath)), 0x001A)
}
func SetTheme() {
}
