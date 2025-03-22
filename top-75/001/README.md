# Two Sum Problem

Hash maps have a constant time look up of O(1). This is great when compared to O(n) using a brute force with two loops.

The downside if the O(n) in memory.

The hash map is a great solution because under the hood it creates a hash for the key with the address of the value, giving a constant time look up of O(1).

## How to solve this problem
- If i do a brute force with two for loops, the complexility would be o(n^2).
- If I need to walk the slice more than once, than the complexity would be O(nˆ2). With hashmaps, the lookup is O(1).
- The trade off is Space Complexity of O(n).
- I can improve this to be O(n) if I only walk the array once and create the elements in the hashmap as I do so.