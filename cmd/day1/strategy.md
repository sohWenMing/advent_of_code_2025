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
- after every iteration, we can just check the number of iterations we need to get through, and at every iteration check if we are already at a left or right endpoint so that we have to know whether or not we have to reset
