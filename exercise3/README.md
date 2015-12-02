# Approach

## Version1

Using an extra array with length 1,000,000 to store the occurances of each element in input array. If current input array element is i, then the entry of index i in the extra array plus 1. After finishing traversing the input array, go through the extra array to check the entry of each index. If it is more than one, mark down the index because it is a duplicate integer. Time complexity: O(n), space complexity: O(n)

## Version2

When traversing the input array, for each element encountered, set its corresponding index value to negative (if current encountered element is 2, set value of array[2] to -value). Because all integers are positive, then next time if encountering another element and its corresponding index value is already negative, we know it is a duplicate. When traversing, we take abs(element) in case it has been flagged to negative. Also this only works safely if the length of array is bigger or equal than 1,000,000 in case of the out of range issue. Time complexity: O(n), space complexity: O(1)