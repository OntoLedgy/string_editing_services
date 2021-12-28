package factories

import (
	"github.com/OntoLedgy/string_editing_services/interfaces"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

func CreateStringEditRange(start_position int, range_length int) interfaces.IStringEditRanges {

	string_edit_range := new(object_model.StringEditRanges)

	string_edit_range.Constructor(start_position, range_length)

	return string_edit_range

}
