package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/midry3125/mouseum/internal/argparse"
	"github.com/midry3125/mouseum/internal/manager"
	"github.com/midry3125/mouseum/internal/utils"
)

func Run() {
	args := argparse.Parse()
	info := manager.Information{}
	info.Target = args.CollectionName
	switch args.Action {
	case argparse.ADD:
		if len(args.References) == 0 {
			utils.RaiseError("`mouse add` needs references.")
		}
		info.Add(args.References)
	case argparse.LIST:
		if !utils.Exists(filepath.Join(manager.ROOT, info.Target)) {
			utils.RaiseError(info.Target + " does not exist.")
		}
		if info.Target == "" {
			fmt.Println("------ Collections ------")
		} else {
			fmt.Println("------ References ------")
		}
		list, err := os.ReadDir(info.GetCollectionPath())
		if err != nil {
			utils.RaiseError("Cannot open: " + info.GetCollectionPath())
		}
		for n, f := range list {
			fmt.Printf("\033[36m%d\033[0m: %s\n", n+1, f.Name())
		}
	case argparse.REMOVE:
		info.RemoveThis()
	case argparse.READ:
		if len(args.References) == 0 {
			utils.RaiseError("`mouse read` needs references.")
		}
		for _, p := range args.References {
			os.Stdout.Write(utils.ReadFile(info.Join(p)))
		}
	case argparse.OPEN:
		if len(args.References) == 0 {
			utils.RaiseError("`mouse open` needs references.")
		}
		for _, p := range args.References {
			utils.OpenFileWithStandard(info.Join(p))
		}
	}
}
