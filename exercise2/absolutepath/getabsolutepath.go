package absolutepath

import (
	"strings"
)

// Take absolute currentPath and a relative target path.
// Return absolute target path.
func GetAbsPath(currentPath, relativePath string) (absPath string) {
	// Special case: currentPath is "" or relativePath is already an absolute path.
	if len(currentPath) == 0 || (len(relativePath) > 0 && string(relativePath[0]) == "/") {
		absPath = GetAbsPath(relativePath, "")
		return
	}

	// Split currentPath and relativePath by "/".
	// Because currentPath start with a "/", get rid of it before split.
	// Otherwise the first string of currentPathArray would be a space.
	if string(currentPath[0]) == "/" {
		currentPath = currentPath[1:]
	}

	currentPathArray := make([]string, 0)
	if len(currentPath) > 0 {
		currentPathArray = strings.Split(currentPath, "/")
	}

	relativePathArray := make([]string, 0)
	if len(relativePath) > 0 {
		relativePathArray = strings.Split(relativePath, "/")
	}

	validPathNames := make([]string, 0)

	validPathNames = helper(currentPathArray, validPathNames)
	validPathNames = helper(relativePathArray, validPathNames)

	// Then, append relativePath part
	for j := range validPathNames {
		absPath += "/"
		absPath += validPathNames[j]
	}
	return
}

func helper(pathArray, validArray []string) []string {
	for i := range pathArray {
		if pathArray[i] == "." {
			continue
		} else if pathArray[i] == ".." {
			if len(validArray) == 0 {
				continue
			} else {
				validArray = validArray[:len(validArray)-1]
			}
		} else {
			validArray = append(validArray, pathArray[i])
		}
	}
	return validArray
}
