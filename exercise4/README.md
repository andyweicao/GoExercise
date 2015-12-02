## Approach

Directly compare the heights of left child node and right child node. If the height difference is more than one, return false. If the difference is equal or smaller than one, going down to check left child node and right child node.

Time complexity for getHeight() is O(n) so the total time complexity is O(n^2). Space complexity is O(1).