## Approach

Use the depth-first search style of recursion. Every time, loop through all valid numbers which starts from the current character and recursion on substrings left. When we have four sets of valid number and the whole input string has been traversed, we append the numbers and store it into result. Time complexity: O(2^n). Space complexity: worst case O(2^n) to store the values.