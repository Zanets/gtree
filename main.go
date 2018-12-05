package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"os/user"
	"os"
	"bufio"
	"strings"
)

// global variables
var regexp_ignore = regexp.MustCompile("ignore$")
var current_user, _ = user.Current()
var irs []*regexp.Regexp 
var repo *Repository = nil


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

// get rules from single directory
func getIgnoreRule(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() && regexp_ignore.MatchString(file.Name()) {
			fp, err := os.Open(path + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer fp.Close()

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				line := string(strings.TrimSpace(scanner.Text()))
				
				// ignore empty line and comment
				if line == "" || strings.HasPrefix(line, "#") {
					continue
				}
				
				// TODO: error check
				reg_ignore := regexp.MustCompile(line)

				irs = append(irs, reg_ignore)
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		} 
	}
}


// entry to get ignore rules
func getIgnoreRules(path string) { 
	
	// get $HOME ignore
	getIgnoreRule(current_user.HomeDir)

	// get ignore files from travelsal
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
		
		for _, ir := range irs {
			if ir.MatchString(name) {
				continue
			}
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

func initRepo(path string) int {
	repo = OpenRepository(path)
	if repo == nil {
		return -1
	}
	return 0
}

func main() {
	target, err := filepath.Abs("./")
	if err != nil {
		log.Fatal(err)
	}
	InitGit()
	if initRepo(target) < 0 {
		os.Exit(1)
	}
	

	// TODO use flag to enable or disable ignore rules
	//getIgnoreRules(target)

	nodes := walk(target, 0)

	for _,node := range nodes {
		fmt.Println(node)
	}
}