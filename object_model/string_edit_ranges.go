package object_model

type StringEditRanges struct {
	Start_position int
	Range_length   int
}

func (string_edit_range *StringEditRanges) Get_start_position() int {

	return 0
}

func (string_edit_range *StringEditRanges) Get_end_position() int {

	return 0
}

func (string_edit_range *StringEditRanges) Constructor(start_position, range_length int) {

	string_edit_range.Start_position = start_position
	string_edit_range.Range_length = range_length
}
