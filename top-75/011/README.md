# 371. Sum of Two Integers

This is tricky problem. Remember the CPU uses binary to add numbers. Therefore I can add numbers using AND and XOR operator remembering I neeed to shift the output of AND operator to the left. 

This will result in a new addition of XOR operation and AND operation with shift (i.e. carry). This needs to go on until there is no more carry (i.e. it is zero).