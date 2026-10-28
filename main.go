//go:build windows

// MacaronCut launcher: writes the embedded editor page to the user's local
// app-data folder and opens it in Edge (or Chrome) as a standalone app window.
// Nothing is uploaded; all video processing happens inside the browser.
package main

import (
	_ "embed"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

//go:embed index.html
var page []byte

func main() {
	dir, err := os.UserCacheDir() // %LocalAppData%
	if err != nil {
		dir = os.TempDir()
	}
	dir = filepath.Join(dir, "MacaronCut")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		fail("無法建立資料夾：" + err.Error())
	}
	file := filepath.Join(dir, "index.html")
	if err := os.WriteFile(file, page, 0o644); err != nil {
		fail("無法寫入剪輯器檔案：" + err.Error())
	}
	u := (&url.URL{Scheme: "file", Path: "/" + filepath.ToSlash(file)}).String()

	for _, b := range browsers() {
		if _, err := os.Stat(b); err == nil {
			cmd := exec.Command(b, "--app="+u, "--window-size=1440,920")
			if cmd.Start() == nil {
				return
			}
		}
	}
	// No Edge/Chrome found: open with the default browser.
	if exec.Command("rundll32", "url.dll,FileProtocolHandler", file).Start() != nil {
		fail("找不到瀏覽器。請手動用 Edge 或 Chrome 開啟：\n" + file)
	}
}

func browsers() []string {
	pf := os.Getenv("ProgramFiles")
	pf86 := os.Getenv("ProgramFiles(x86)")
	local := os.Getenv("LocalAppData")
	return []string{
		filepath.Join(pf86, `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(pf, `Microsoft\Edge\Application\msedge.exe`),
		filepath.Join(pf, `Google\Chrome\Application\chrome.exe`),
		filepath.Join(pf86, `Google\Chrome\Application\chrome.exe`),
		filepath.Join(local, `Google\Chrome\Application\chrome.exe`),
	}
}

func fail(msg string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	box := user32.NewProc("MessageBoxW")
	t, _ := syscall.UTF16PtrFromString(msg)
	c, _ := syscall.UTF16PtrFromString("馬卡龍剪輯室")
	box.Call(0, uintptr(unsafe.Pointer(t)), uintptr(unsafe.Pointer(c)), 0x10)
	os.Exit(1)
}
