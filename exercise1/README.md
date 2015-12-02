# Approach

Two queues, queue1 and queue2, are used to work as a stack. Space complexity is O(n) because we need to store each element.

## Push

When pushing a new element, always offer it into queue1. Time complexity: O(1)

## Pop

If queue1 only has one element, poll it out. If queue1 is more than one element, poll elements from queue1 to queue2 until last element. Poll it out. When finishing pop, switch names of queue2 and queue1.(Always make sure queue2 is empty for next operation.) Worst time complexity for each pop operation: O(n)

