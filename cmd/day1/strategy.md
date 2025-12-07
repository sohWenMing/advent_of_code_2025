# Key Points

- the dial starts at 50

* if the dial is already at 99, moving it right makes it restart at 0
* if the dial is at 0, moving it left makes it go to 99

## Reading from the input

- the simplest way to get all the inputs one by one is to just read the file inline by line using something like scanline
- since the lines will always be prefixed with "R" or "L", we can always just gget the prefix first character to know whether it's right or left

## Managing input after it is read

- strip the first letter as a prefix, so we can know whether the overall position is left or right
- parseInt the number, so we can get the number of iterations
- after each set of rotations is done, check if the current position % 100 == 0, which would signify that it was
- a number that signified 0

## Part 2

- The main idea is to keep adding to the checked substring which is the first half of the string
- keep using trimPrefix, and then check index to see if it's found at index 0, if it is, then keep going
- if the string being checked in the moment is finally empty, then it is invalid
