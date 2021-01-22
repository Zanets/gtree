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
		for i := 0 ; i < subnode.Level; i++ {
			fmt.Print("  ")
		}

		fmt.Println(subnode.Name)
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

	Repo.Close()
}