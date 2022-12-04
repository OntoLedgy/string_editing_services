package factories

import (
	"github.com/OntoLedgy/string_editing_services/interfaces"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type StringEditorFactory struct{}

func (factory *StringEditorFactory) CreateStringEditor(
	stringToEdit string) interfaces.IStringEditors {

	stringEditor := new(
		object_model.StringEditors)

	stringEditor.
		Constructor(
			stringToEdit)

	return stringEditor
}
