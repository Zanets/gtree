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

func walk(path string, level int, nodes []Node) []Node {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := file.Name()
		if name == ".git" {
			continue
		}
		if file.IsDir() {
			nodes = append(nodes, Node{file.Name(), path, Directory, level})
			nodes = walk(path + "/" + file.Name(), level+1, nodes)
		} else {
			nodes = append(nodes, Node{file.Name(), path, File, level})
		}
	}
	return nodes
}

func main() {
	target, err := filepath.Abs("./")
	var nodes []Node
	if err == nil {
		nodes = walk(target, 0, nodes)
		for _,node := range(nodes) {
			fmt.Println(node)
		}
	} else {
		log.Fatal(err)
	}
}