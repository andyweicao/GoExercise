package main

import (
	"./duplicate"
	"fmt"
)

func main() {
	array := []int{1, 3, 4, 2, 2, 3, 6, 7, 8, 12, 7}
	duplicates := duplicate.FindDuplicates(array)
	fmt.Printf("Here are all duplicates: %v\n", duplicates)
}
