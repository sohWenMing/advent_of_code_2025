package day5internal

func GetPart2(filepath string) (numIngredients int64, err error) {
	_, startEnds, err := GetNumsAndStartEndsFromFile(filepath)
	if err != nil {
		return 0, err
	}
	numIngredients = 0
	workingStartEnds := []StartEnd{}
	for _, startEnd := range startEnds {
		workingStartEnds = RecursiveAppendStartEnd(startEnd, workingStartEnds)
	}
	for _, workingStartEnd := range workingStartEnds {
		numIngredients += workingStartEnd.End - workingStartEnd.Start + 1
	}
	return numIngredients, nil
}

func RecursiveAppendStartEnd(inStartEnd StartEnd, inSlice []StartEnd) (result []StartEnd) {
	// empty slice, just append
	if len(inSlice) == 0 {
		result := []StartEnd{inStartEnd}
		return result
	}

	current := inSlice[0]

	// start and end are within first index, startEnd is already accounted for, just return
	if inStartEnd.Start >= current.Start && inStartEnd.End <= current.End {
		return inSlice
	}

	// start end and are both before the start of the first index, add inStartEnd to beginning
	// of resulting slice, then append the whole input slice
	if inStartEnd.End < current.Start {
		result := []StartEnd{inStartEnd}
		result = append(result, inSlice...)
		return result
	}

	remainingSlice := getRemainingSlice(inSlice)

	// if the startEnd starts after current, just add current to start, then append result of calling function
	// with unaltered startend
	if inStartEnd.Start > current.End {
		returned := []StartEnd{current}
		returnedFromFunc := RecursiveAppendStartEnd(inStartEnd, remainingSlice)
		returned = append(returned, returnedFromFunc...)
		return returned
	}

	// at this point, inStartEnd.Start has to be <= current.End
	returned := []StartEnd{}

	// if start of inStartEnd, is before current start, eval if need to amend and append
	if inStartEnd.Start < current.Start {
		workingStartEnd := StartEnd{inStartEnd.Start, current.Start - 1}
		if evalStartEndValid(workingStartEnd) {
			returned = append(returned, workingStartEnd)
		}
	}

	// either way, append current
	returned = append(returned, current)

	// check the value of the working startend, assuming that the new start is current end + 1
	workingStartEnd := StartEnd{current.End + 1, inStartEnd.End}

	// if the new start end is valid, then call teh recursive function using the new working StartEnd, and the
	// remaining slice
	if evalStartEndValid(workingStartEnd) {
		returnedFromFunc := RecursiveAppendStartEnd(workingStartEnd, remainingSlice)
		returned = append(returned, returnedFromFunc...)
		return returned
		// else, just append the remaining slice to the return value, which should already have current at the start
	} else {
		returned = append(returned, remainingSlice...)
		return returned
	}
}

func getRemainingSlice(inSlice []StartEnd) []StartEnd {
	var remainingSlice []StartEnd
	if len(inSlice) > 1 {
		remainingSlice = inSlice[1:]
	} else {
		remainingSlice = []StartEnd{}
	}
	return remainingSlice
}

func evalStartEndValid(startEnd StartEnd) bool {
	return startEnd.End >= startEnd.Start
}
