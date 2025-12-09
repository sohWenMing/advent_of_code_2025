# Removing as many possible crates using the forklift

- at this point, we've already to establish within the first pass how many crates are removable using a forklift
- to remember - a crate can only be removed if it's surrounded by less than 4 crates

## Noting that a crate has been removed

- since cells that exist are kept within the map - what we can do is that for each crate that is removed, we can delete it from the map at the point where we have decided that it's already removable

* we can leverage on the first initial function

```

func GetCellsValsFromFile(filePath string) (cellmap FilledCellMap, allCells []Cell, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	cellmap = make(FilledCellMap)
	allCells = []Cell{}

	scanner := bufio.NewScanner(file)
	yIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		readToFilledCellMap(cellmap, line, yIdx)
		readToAllCells(&allCells, line, yIdx)
		yIdx++
	}
	return cellmap, allCells, nil
}
```

this will give us the starting state of:

- the initial map - which will note which cells have crates
- the layout of the cells
