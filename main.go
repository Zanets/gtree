package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type NodeType int


const (
	File NodeType = 0
	Directory NodeType = 1
)

type Node struct {
	Name string 
	Path string
	Type NodeType
	Level int
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

func walk(path string, level int) []Node {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var nodes []Node
	for _, file := range files {
		// drawBranch(file.Name(), level)
		name := file.Name()
		if name == ".git" {
			continue
		}
		if file.IsDir() {
			nodes = append(nodes, walk(path + "/" + file.Name(), level+1)...)
		} else {
			nodes = append(nodes, Node{file.Name(), path, File, level})
		}
	}
	return nodes
}

func main() {
	target, err := filepath.Abs("./")
	fmt.Println(target)
	if err == nil {
		nodes := walk(target, 1)
		fmt.Println(nodes)
	} else {
		log.Fatal(err)
	}
}