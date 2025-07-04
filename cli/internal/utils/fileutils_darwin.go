package utils

import (
	"os"
	"path/filepath"
)

func OpenFileWithStandard(path string) {
	if !Exists(path) {
		RaiseError(filepath.Base(path) + " does not exist.")
	}
	os.StartProcess("open", []string{path}, &os.ProcAttr{
		Setpgid: true,
	})
}
