package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	printFiles("/home/salvo")
}

func printFiles(dirpath string) {
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fullname := filepath.Join(dirpath, f.Name())
		if f.IsDir() {
			printFiles(fullname)
		} else {
			fmt.Println(fullname)
		}
	}
}
