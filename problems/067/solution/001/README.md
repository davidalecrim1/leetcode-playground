# 067. Add Binary
https://leetcode.com/problems/add-binary/

## Example of the Solution

Let’s say we are adding two binary numbers, a = "1010" and b = "1101", and carry = 0.

Step 1: Start with the least significant bits:
	•	a[3] = '0', b[3] = '1', carry = 0.
	•	sum = 0 + 1 + 0 = 1.
	•	sum % 2 = 1 → Current bit = '1'.
	•	sum / 2 = 0 → No carry.
	•	result = "1".

Step 2: Move to the next bit:
	•	a[2] = '1', b[2] = '0', carry = 0.
	•	sum = 1 + 0 + 0 = 1.
	•	sum % 2 = 1 → Current bit = '1'.
	•	sum / 2 = 0 → No carry.
	•	result = "11".

Step 3: Move to the next bit:
	•	a[1] = '0', b[1] = '1', carry = 0.
	•	sum = 0 + 1 + 0 = 1.
	•	sum % 2 = 1 → Current bit = '1'.
	•	sum / 2 = 0 → No carry.
	•	result = "111".

Step 4: Move to the most significant bit:
	•	a[0] = '1', b[0] = '1', carry = 0.
	•	sum = 1 + 1 + 0 = 2.
	•	sum % 2 = 0 → Current bit = '0'.
	•	sum / 2 = 1 → Carry = 1.
	•	result = "0111".

Step 5: Handle remaining carry:
	•	Since carry = 1, prepend '1' to the result.
	•	result = "10111".