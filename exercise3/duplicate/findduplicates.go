package duplicate

// Since integers range from 1 to 1,000,000, use an array with length of 1,000,000 to
// record occurances of each integer.
// Then, go through record array and insert integers with orrcuance more than 1 into result.
// Time complexity is O(n).
// When the array is large enough compared to 1,000,000, the space complexity can be considered as constant(O(1)).
func FindDuplicates(input []int) (result []int) {
	if len(input) < 2 {
		return
	}
	record := [1000000]int{}
	for i := range input {
		record[input[i]-1] += 1
	}

	for j := range record {
		if record[j] > 1 {
			result = append(result, j+1)
		}
	}
	return
}

func FindDuplicatesVersion2(input []int) (result []int) {
	dups := make(map[int]bool)
	for i := range input {
		var current int
		if input[i] < 0 {
			current = -1 * input[i]
		} else {
			current = input[i]
		}

		if input[current] < 0 {
			dups[current] = true
		} else {
			input[current] *= -1
		}
	}

	for i := range dups {
		result = append(result, i)
	}
	return
}
