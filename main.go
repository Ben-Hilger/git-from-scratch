package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args[1:]

	numberOfArgs := len(argsWithProg)

	if numberOfArgs == 0 {
		fmt.Println("Warning, you'll need to enter one of the commands: init, status, add, commit, push")
		return
	}

	parseCommand(argsWithProg)
}

func parseCommand(args []string) {
	if len(args) == 0 {
		return
	}

	if args[0] == "init" {
		initGitProject()
	}
}

func initGitProject() {
	file, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if hasGitFileInDirectory(file) {
		fmt.Println("A git project already exists in the current folder")
		return
	}
	fmt.Println("Initializing git project in the current folder")
	err = os.Mkdir(fmt.Sprintf("%s/.git", file), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("New folder created")
}

func hasGitFileInDirectory(file string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/.git", file))
	return !os.IsNotExist(err)
}
