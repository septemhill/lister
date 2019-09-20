package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

var listCount int64

func listLimitCheck() bool {
	if listLimit == maxListLimit || listCount < listLimit {
		listCount++
		return true
	}

	return false
}

func output(ftype FileType, path string) {
	if !listLimitCheck() {
		return
	}

	if ftype == TYPE_FILE {
		fmt.Println(FileTag, path)
	} else {
		fmt.Println(DirectoryTag, path)
	}
}

func listFileWithFullPath(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	for _, file := range files {
		if file.IsDir() {
			output(TYPE_DICT, path+"/"+file.Name())
			listFileWithFullPath(path + "/" + file.Name())
		} else {
			output(TYPE_FILE, path+"/"+file.Name())
		}
	}
}

func mainAction(c *cli.Context) {
	path := c.String("dir")
	listFileWithFullPath(path)
	fmt.Printf("Total: %d\n", listCount)
}

func main() {
	app := cli.NewApp()
	app.Usage = "List files with full path"
	app.Author = "Septem Li"
	app.Action = mainAction
	app.Version = "0.0.2"
	app.Flags = []cli.Flag{
		DirFlag,
		FindFlag,
		CaseSensitiveFlag,
		ListLimitFlag,
		TypeDisplayFlag,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("WTF?")
	}

	fmt.Println(caseSensitive)
}
