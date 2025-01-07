# Neet Code Blind 75 List

The goal is to use this problem list to Pareto the best problems to learn the most about data structures and algorithms.

## Notes

| ID  | Name | Difficulty                                                                                                         | Notes                                                                                                                                                                         |
| --- | ---- | ------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 001 | Easy | [1. Two Sum Problem](https://leetcode.com/problems/two-sum/)                                                       | Do a one time pass using a hashmap for previous passed elements. Don't check the same element twice.                                                                          |
| 002 | Easy | [121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/) | Sliding window with two pointers for low and high.                                                                                                                            |
| 003 | Easy | [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/description/)                           | Create a hash map interating through the array. Validate if posterior values exists in it.                                                                                    |
| 005 | Easy | [53. Maximum Subarray ](https://leetcode.com/problems/maximum-subarray/description/)                               | Use a "sliding window" and check if the curent sum if below zero to reset to throw out the numbers.                                                                           |
| 010 | Easy | [371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/description/)                         | Use binary sum with XOR and AND with shift to the left.                                                                                                                       |
| 011 | Easy | [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/description/)                               | parse to base 2 (go: strconv) and use count of values.                                                                                                                        |
| 013 | Easy | [268. Missing Number](https://leetcode.com/problems/missing-number/description/)                                   | solution 1: calculate the sum of the array minus the given array using a variable with len(nums); solution 2: sort the array and check with the index (but it is O(n log n)). |
| 014 | Easy | [190. Reverse Bits](https://leetcode.com/problems/reverse-bits/description/)                                       | Format the input as binary string and reverse the string with a loop.                                                                                                         |




## Reference List
- [Neetcode Blind 75 Questions List](https://docs.google.com/spreadsheets/d/1A2PaQKcdwO_lwxz9bAnxXnIQayCouZP6d-ENrBz_NXc/edit?gid=0#gid=0)
- [Leetcode Top 150 Questions](https://leetcode.com/studyplan/top-interview-150/)