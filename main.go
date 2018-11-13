package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Node struct {
	Name string 
}

func isGit(path string) bool {
	return false	
}

func drawBranch(name string, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	fmt.Println("|-" + name)
}

func walk(path string, level int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	
	for _, file := range files {
		drawBranch(file.Name(), level)
		if file.IsDir() {
			walk(path + "/" + file.Name(), level+1)
		}
	}
}

func main() {
	target, err := filepath.Abs("./")
	fmt.Println(target)
	if err == nil {
		walk(target, 1)
	} else {
		log.Fatal(err)
	}
}