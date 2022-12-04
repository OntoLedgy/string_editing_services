package interfaces

import (
	coreObjectModelInterfaces "github.com/OntoLedgy/core_ontology/code/core/contracts"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type IStringEditors interface {
	coreObjectModelInterfaces.IObjects

	Terminate() *object_model.StringEditHistories

	Insert(
		editRange object_model.StringEditRanges,
		stringToInsert string)

	Delete(
		editRanges object_model.StringEditRanges)

	Substitute(
		editRange object_model.StringEditRanges,
		stringToSubstitute string)

	GetCurrentString() object_model.CurrentStrings

	GetStringEditHistory() *object_model.StringEditHistories
}
