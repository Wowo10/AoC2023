package main

type RangeCollection struct {
	MyRanges []MyRange
}

func (c RangeCollection) GetDestination(i int) int {
	for _, r := range c.MyRanges {
		if r.IsInRange(i) {
			return r.GetDestination(i)
		}
	}

	return i
}

type MyRange struct {
	destination int
	source      int
	length      int
}

func CreateMyRange(destination, source, length int) MyRange {
	return MyRange{
		destination: destination,
		source:      source,
		length:      length,
	}
}

func (r MyRange) IsInRange(i int) bool {
	return i-r.source >= 0 && i-r.source < r.length
}

func (r MyRange) GetDestination(i int) int {
	return r.destination + i - r.source
}
