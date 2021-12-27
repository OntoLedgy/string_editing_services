package object_model

import (
	"core_foundation/core_object_model"
	"fmt"
)

type StringEditors struct {
	core_object_model.Objects
	CurrentString       string
	String_to_edit      string
	String_edit_history *StringEditHistories
}

func (string_editor *StringEditors) Constructor(string_to_edit string) {

	//#TODO this should not need to be assigned explicitly, should just be a setter.
	string_editor.Object_uuid = string_editor.Set_object_uuid()

	string_editor.String_to_edit =
		string_to_edit

	string_editor.CurrentString =
		string_to_edit

	string_editor.
		String_edit_history =
		new(StringEditHistories)

	string_editor.String_edit_history.
		Create(string_to_edit)
}

func (
	string_editor *StringEditors) Terminate() *StringEditHistories {

	return string_editor.String_edit_history

}

func (
	string_editor *StringEditors) Insert(
	edit_range StringEditRanges,
	string_to_insert string) {
	fmt.Print("inserting\n")
}

func (
	string_editor *StringEditors) Delete(
	edit_range StringEditRanges) {

}

func (
	string_editor *StringEditors) Substitute(
	edit_range StringEditRanges,
	string_to_substitute string) {

}

func (
	string_editor *StringEditors) GetCurrentString() CurrentStrings {

	var current_string CurrentStrings

	return current_string
}

func (
	string_editor *StringEditors) Get_string_edit_history() *StringEditHistories {

	return string_editor.String_edit_history
}
