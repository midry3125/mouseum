package manager

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/midry3125/mouseum/internal/utils"
)

var (
	ROOT string = GetRootDir()
)

type Information struct {
	Target string
}

func (i *Information) GetCollectionPath() string {
	res := filepath.Join(ROOT, i.Target)
	if !utils.Exists(res) {
		os.Mkdir(res, os.ModePerm)
	}
	return res
}

func (i *Information) Join(name string) string {
	return filepath.Join(i.GetCollectionPath(), name)
}

func (i *Information) Add(refs []string) {
	for _, p := range refs {
		if utils.IsDir(p) {
			f, _ := os.ReadDir(p)
			var r []string
			for _, q := range f {
				r = append(r, filepath.Join(p, q.Name()))
			}
			i.Add(r)
		} else {
			f, err := os.Open(p)
			if err != nil {
				utils.RaiseError("Cannot open as file: " + p)
			}
			defer f.Close()
			dst := i.Join(filepath.Base(p))
			d, err := os.Create(dst)
			if err != nil {
				utils.RaiseError("Cannot open as file: " + dst)
			}
			defer d.Close()
			_, err = io.Copy(d, f)
			if err != nil {
				utils.RaiseError("Faild to copy: " + p)
			}
		}
	}
}

func (i *Information) RemoveThis() {
	os.RemoveAll(i.GetCollectionPath())
}

func (i *Information) Remove(name string) {
	os.Remove(i.Join(name))
}

func GetRootDir() string {
	home, _ := os.UserHomeDir()
	var appdir string
	switch runtime.GOOS {
	case "windows":
		appdir = os.Getenv("APPDATA")
		if appdir == "" {
			appdir = filepath.Join(home, "AppData", "Roaming")
		}
	case "darwin":
		appdir = filepath.Join(home, "Library", "Application Support")
	default:
		appdir = filepath.Join(home, ".local", "share")
	}
	p := filepath.Join(appdir, "mouseum")
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		os.Mkdir(p, os.ModePerm)
	}
	return p
}
