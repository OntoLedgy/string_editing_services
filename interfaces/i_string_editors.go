package interfaces

import (
	core_object_model_interfaces "github.com/OntoLedgy/core_ontology/interfaces"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type IStringEditors interface {
	core_object_model_interfaces.IObjects

	Terminate() *object_model.StringEditHistories

	Insert(
		edit_range object_model.StringEditRanges,
		string_to_insert string)

	Delete(
		edit_range object_model.StringEditRanges)

	Substitute(
		edit_range object_model.StringEditRanges,
		string_to_substitute string)

	GetCurrentString() object_model.CurrentStrings

	Get_string_edit_history() *object_model.StringEditHistories
}
