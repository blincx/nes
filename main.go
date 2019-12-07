package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"fmt"
	"github.com/fogleman/nes/ui"
)

func main() {
	log.SetFlags(0)
	fmt.Println("A0")
	paths := getPaths()
	fmt.Println("A1")
	if len(paths) == 0 {
		log.Fatalln("no rom files specified or found")
	}
	fmt.Println("A2")
	ui.Run(paths)
}

func getPaths() []string {
	var arg string
	args := os.Args[1:]
	if len(args) == 1 {
		arg = args[0]
	} else {
		arg, _ = os.Getwd()
	}
	info, err := os.Stat(arg)
	if err != nil {
		return nil
	}
	fmt.Println("A3")
	if info.IsDir() {
		infos, err := ioutil.ReadDir(arg)
		if err != nil {
			return nil
		}
		var result []string
		for _, info := range infos {
			name := info.Name()
			if !strings.HasSuffix(name, ".nes") {
				continue
			}
			result = append(result, path.Join(arg, name))
		}
		return result
	} else {
		return []string{arg}
	}
}
