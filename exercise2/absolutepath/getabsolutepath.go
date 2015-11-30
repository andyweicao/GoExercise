package absolutepath

import (
	"strings"
)

// Take absolute currentPath and a relative target path.
// Return absolute target path.
func GetAbsPath(currentPath, relativePath string) (absPath string) {
	// Special case: relativePath is "".
	if len(relativePath) == 0 {
		absPath = currentPath
		return
	}
	// Special case: currentPath is "" or relativePath is already an absolute path.
	if len(currentPath) == 0 || string(relativePath[0]) == "/" {
		absPath = relativePath
		return
	}

	// Split currentPath and relativePath by "/".
	// Because currentPath start with a "/", get rid of it before split.
	// Otherwise the first string of currentPathArray would be a space.
	if string(currentPath[0]) == "/" {
		currentPath = currentPath[1:]
	}
	currentPathArray := strings.Split(currentPath, "/")
	currentPathArrayLength := len(currentPathArray)
	relativePathArray := strings.Split(relativePath, "/")
	// prevCount will count how many "../" in relativePath
	// stopPostion will represent the index of the first shown directory name in relativePathArray
	var prevCount int
	var stopPosition int
	for i := range relativePathArray {
		if relativePathArray[i] == ".." {
			prevCount++
		} else if relativePathArray[i] == "." {
			continue
		} else {
			stopPosition = i
			break
		}
	}

	// Discard invalid "../" if there are too many of it
	if prevCount >= currentPathArrayLength {
		prevCount = currentPathArrayLength
	}

	// First, append currentPath part
	for i := 0; i < currentPathArrayLength-prevCount; i++ {
		absPath += "/"
		absPath += currentPathArray[i]
	}

	// Then, append relativePath part
	for j := stopPosition; j < len(relativePathArray); j++ {
		absPath += "/"
		absPath += relativePathArray[j]
	}
	return
}
