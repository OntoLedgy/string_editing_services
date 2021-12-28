package factories

import (
	"github.com/OntoLedgy/string_editing_services/interfaces"
	"github.com/OntoLedgy/string_editing_services/object_model"
)

type StringEditorFactory struct{}

func (factory *StringEditorFactory) CreateStringEditor(string_to_edit string) interfaces.IStringEditors {

	string_editor := new(
		object_model.StringEditors)

	string_editor.
		Constructor(
			string_to_edit)

	return string_editor
}
