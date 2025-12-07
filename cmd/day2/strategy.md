# Getting all subsets of the string

## Dealing with the string after parsing the number to a string

- a number is invalid if it is made up of on a sequence of numbers that is repeated
- the repeating of the characters has to take up the whole string

### Strategy

The overall strategy involves reading in the file and the splitting by ","
After that we can just get the range of the numbers by splitting each line by "-", and then parsing the stringified
numbers to integers and getting the range between the two numbers (inclusive)
Split each number, if the first set of characters matches the last set, the it's invalid
Add up all invalid numbers
