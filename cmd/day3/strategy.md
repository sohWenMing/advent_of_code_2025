# Finding the largest possible joltage

- a joltage is made up of a combination of 2 characters, and position matters
- first, find the largest possible number and track its position. we only want to replace the number if the following
  numbers if larger, because if not having the number in an earlier position gives better chance to get higher numbers
- after that, substring from after the position, to get the rest of the characters
- find the largest number in the resulting substring
