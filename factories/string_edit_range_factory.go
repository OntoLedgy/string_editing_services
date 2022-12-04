package factories

import (
	"github.com/OntoLedgy/string_editing_services/interfaces"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

func CreateStringEditRange(startPosition int, rangeLength int) interfaces.IStringEditRanges {

	stringEditRange := new(object_model.StringEditRanges)

	stringEditRange.Constructor(startPosition, rangeLength)

	return stringEditRange

}
