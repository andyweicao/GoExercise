package main

import (
	"./absolutepath"
	"fmt"
)

func main() {
	currentDirectoryPath := "/home/andy"
	targetPath := "./bin/foo"
	absTargetPath := absolutepath.GetAbsPath(currentDirectoryPath, targetPath)
	fmt.Println(absTargetPath)
}
