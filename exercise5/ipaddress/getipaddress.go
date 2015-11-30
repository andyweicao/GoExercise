package ipaddress

import (
	"strconv"
)

// Use a helper function to find all possible combinations recursively
func GetIpAddress(s string) (result []string) {
	list := make([]string, 0)
	if len(s) < 4 || len(s) > 12 {
		return
	}
	helper(&result, &list, s, 0)
	return
}

// list is used to store current valid numbers
// result is used to store valid combinations
func helper(result, list *[]string, s string, start int) {
	if len(*list) == 4 {
		if start != len(s) {
			return
		}
		// Have four valid numbers in list now and no more number in s
		// Restore ip by combining these four numbers
		var current string
		for i := range *list {
			current += (*list)[i]
			current += "."
		}
		*result = append(*result, current[0:(len(current)-1)])
		return
	}

	// Find all valid numbers starting at start
	for i := start; i < len(s) && i < start+3; i++ {
		temp := s[start : i+1]
		if isvalid(temp) {
			*list = append(*list, temp)
			helper(result, list, s, i+1)
			*list = (*list)[0:(len(*list) - 1)]
		}
	}
}

// check whether s is in range 0-255
func isvalid(s string) bool {
	if string(s[0]) == "0" {
		return s == "0" // Discard "00" and "000"
	}
	value, _ := strconv.ParseInt(s, 10, 0)
	return value >= 0 && value <= 255
}
