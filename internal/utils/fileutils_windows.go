package utils

import (
	"os/exec"
	"path/filepath"
	"syscall"
)

func OpenFileWithStandard(path string) {
	if !Exists(path) {
		RaiseError(filepath.Base(path) + " does not exist.")
	}
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", path)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}
	cmd.Start()
}
