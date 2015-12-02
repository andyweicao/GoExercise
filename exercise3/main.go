package main

import (
	"./duplicate"
	"fmt"
)

func main() {
	array1 := []int{1, 3, 4, 2, 2, 3, 6, 7, 8, 12, 7}
	duplicates1 := duplicate.FindDuplicates(array1)
	fmt.Printf("Here are all duplicates for array1: %v\n", duplicates1)

	array2 := []int{3, 7, 6, 6, 6, 3, 2, 2, 2, 6, 7, 3, 10, 10}
	duplicates2 := duplicate.FindDuplicatesVersion2(array2)
	fmt.Printf("Here are all duplicates for array2: %v\n", duplicates2)
}
