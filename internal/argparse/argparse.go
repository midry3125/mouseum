package argparse

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/midry3125/mouseum/internal/manager"
	"github.com/midry3125/mouseum/internal/utils"
	"golang.org/x/exp/slices"
)

const (
	ADD    string = "add"
	USE    string = "use"
	HELP   string = "help"
	LIST   string = "list"
	REMOVE string = "rm"
	READ   string = "read"
	OPEN   string = "open"
)

var (
	KEYWORDS []string = []string{ADD, USE, HELP, LIST, REMOVE, READ, OPEN}
)

type Arguments struct {
	Action         string
	CollectionName string
	References     []string
}

func Parse() Arguments {
	res := Arguments{}
	args := os.Args[1:]
	if len(os.Args) == 1 {
		utils.ShowHelp()
	} else {
		var (
			v    string
			refs []string
		)
		v, args = GetArg(args)
		if slices.Contains(KEYWORDS, v) {
			res.Action = v
			switch res.Action {
			case HELP:
				utils.ShowHelp()
			default:
				if len(args) == 0 && res.Action != LIST {
					utils.RaiseError("Missing collection name.")
				} else if 0 < len(args) {
					v, args = GetArg(args)
					collections := utils.GetDirsWithBase(manager.ROOT)
					i, err := strconv.Atoi(v)
					if err == nil && 0 < i && i <= len(collections) {
						res.CollectionName = collections[i-1]
					} else {
						res.CollectionName = v
					}
				}
			}
		} else {
			utils.RaiseError("Unknown keyword: " + v)
		}
		if res.Action == OPEN || res.Action == READ {
			refs = utils.GetFilesWithBase(filepath.Join(manager.ROOT, res.CollectionName))
		}
		for 0 < len(args) {
			v, args = GetArg(args)
			switch v {
			// オプションなど
			default:
				n, err := strconv.Atoi(v)
				if err == nil && 0 < n && n <= len(refs) {
					res.References = append(res.References, refs[n-1])
				} else {
					res.References = append(res.References, v)
				}
			}
		}
	}
	return res
}

func GetArg(args []string) (string, []string) {
	res := args[0]
	args = args[1:]
	return res, args
}
