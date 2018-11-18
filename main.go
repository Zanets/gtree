package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"os/user"
)

// global variables
var regexp_ignore = regexp.MustCompile(".*ignore")
var current_user, _ = user.Current()


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

// used to get rules from single directory
func getIgnoreRule(path string) []string {
	var irs []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if regexp_ignore.MatchString(file.Name()) {
		} 
	}
	return irs
}


// entry to get ignore rules
func getIgnoreRules(path string) []string { 
	var irs []string
	// get $HOME ignore
	fmt.Println(current_user)
	// get ignore files from travelsal

	return irs
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
		name := file.Name()
		if name == ".git" {
			continue
		}
		if file.IsDir() {
			nodes = append(nodes, Node{file.Name(), path, Directory, level})
			nodes = append(nodes, walk(path + "/" + file.Name(), level+1)...)
		} else {
			nodes = append(nodes, Node{file.Name(), path, File, level})
		}
	}
	return nodes
}

func main() {
	target, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}


	getIgnoreRules(target)
	nodes := walk(target, 0)

	for _,node := range nodes {
		fmt.Println(node)
	}
}