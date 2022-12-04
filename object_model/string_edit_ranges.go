package object_model

type StringEditRanges struct {
	StartPosition int
	RangeLength   int
}

func (stringEditRange *StringEditRanges) GetStartPosition() int {

	return 0
}

func (stringEditRange *StringEditRanges) GetEndPosition() int {

	return 0
}

func (stringEditRange *StringEditRanges) Constructor(
	startPosition,
	rangeLength int) {

	stringEditRange.StartPosition = startPosition
	stringEditRange.RangeLength = rangeLength
}
