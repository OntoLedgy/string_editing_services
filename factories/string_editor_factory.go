package factories

import (
	"string_editing_services/interfaces"
	"string_editing_services/object_model"
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
