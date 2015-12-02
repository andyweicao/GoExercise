package main

import (
	"./absolutepath"
	"fmt"
)

func main() {
	currentDirectoryPath := "/home/andy"
	targetPath := make([]string, 0)
	targetPath = append(targetPath, "./bin/foo")
	targetPath = append(targetPath, "/home/../")
	targetPath = append(targetPath, "/home/./")
	targetPath = append(targetPath, "/home/andy/../")
	targetPath = append(targetPath, "/home/andy/./")

	for i := range targetPath {
		absTargetPath := absolutepath.GetAbsPath(currentDirectoryPath, targetPath[i])
		fmt.Println(absTargetPath)
	}
}
