package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"path"
	"os"
)


type NodeType int
const (
	NTF NodeType = 0
	NTD NodeType = 1
)

type Node struct {
	Order int // order number in folder
	Name string // file name
	Path string // parent folder path
	Type NodeType
	Level int
	SubFiles []*Node
}

var Repo Repository
var Root Node

func scanNode(node* Node) {
	if node.Type != NTD {
		return 
	}
		
	curPath := node.Path + "/" + node.Name

	subfiles, err := ioutil.ReadDir(curPath)
	if err != nil {
		log.Fatal(err)
	}
	
	for i, subfile := range subfiles {
		
		if subfile.Name() == ".git" {
			continue
		}
		
		isIgnored, _ := Repo.IsIgnored(curPath + "/" + subfile.Name())
		if isIgnored {
			continue
		}

		subnode := Node {
			Order: i, 
			Name: subfile.Name(), 
			Path: curPath, 
			Type: NTF,
			Level: node.Level+1 }
		
		if subfile.IsDir() {
			subnode.Type = NTD
			scanNode(&subnode)
		}
		 
		node.SubFiles = append(node.SubFiles, &subnode)
	}

}

func printNode(node* Node) {
	for _, child := range node.SubFiles {
		for i := 0 ; i < child.Level; i++ {
			fmt.Print("  ")
		}

		fmt.Println(child.Name)
		printNode(child)
	}
}

func main() {
	repoPath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	Repo = Repository{}
	if Repo.Open(repoPath) < 0 {
		fmt.Println("Open repo fail")
		return
	}

	Root = Node{
		Order: 0, 
		Name: path.Base(repoPath), 
		Path: path.Dir(repoPath), 
		Type: NTD, 
		Level: 0}
	
	scanNode(&Root)
	printNode(&Root)


	Repo.Close()
}