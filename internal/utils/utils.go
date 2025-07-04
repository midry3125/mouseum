package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func RaiseError(msg string) {
	fmt.Printf("Error: %s\n", msg)
	os.Exit(1)
}

func ShowHelp() {
	fmt.Println(`Usage:
	$ mouse [action] (options) [references... ]

Actions:
	add [collection name]
		Add new references to specified collection name.

	rm [collection name]
		Remove specified collection name.

	use [collection name]
		Run TUI screen with opening specified collection name.

	list (collection name)
		Show all collection names or references.

	help
		Show this help message.`)
	os.Exit(0)
}

func ReadFile(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		RaiseError("Cannot open: " + path)
	}
	return b
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	i, err := os.Stat(path)
	if err != nil {
		return false
	}
	return i.IsDir()
}

func GetFilesWithBase(path string) []string {
	res := []string{}
	d, err := os.ReadDir(path)
	if err != nil {
		RaiseError(filepath.Base(path) + " does not exists.")
	}
	for _, p := range d {
		f := filepath.Join(path, p.Name())
		if !IsDir(f) {
			res = append(res, p.Name())
		}
	}
	sort.Slice(res, func(i int, j int) bool {
		return res[i] < res[j]
	})
	return res
}

func GetDirsWithBase(path string) []string {
	res := []string{}
	d, err := os.ReadDir(path)
	if err != nil {
		RaiseError(filepath.Base(path) + " does not exists.")
	}
	for _, p := range d {
		f := filepath.Join(path, p.Name())
		if IsDir(f) {
			res = append(res, p.Name())
		}
	}
	sort.Slice(res, func(i int, j int) bool {
		return res[i] < res[j]
	})
	return res
}
