package object_model

import (
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
)

type StringEditors struct {
	objects.BnogObjects
	CurrentString     string
	StringToEdit      string
	StringEditHistory *StringEditHistories
}

func (stringEditor *StringEditors) Constructor(stringToEdit string) {

	//#TODO this should not need to be assigned explicitly, should just be a setter.
	stringEditor.Object_uuid = stringEditor.Set_object_uuid()

	stringEditor.StringToEdit =
		stringToEdit

	stringEditor.CurrentString =
		stringToEdit

	stringEditor.
		StringEditHistory =
		new(StringEditHistories)

	stringEditor.StringEditHistory.
		Create(stringToEdit)
}

func (
	stringEditor *StringEditors) Terminate() *StringEditHistories {

	return stringEditor.StringEditHistory

}

func (
	stringEditor *StringEditors) Insert(
	editRange StringEditRanges,
	stringToInsert string) {

	fmt.Print("inserting\n")

}

func (
	stringEditor *StringEditors) Delete(
	editRange StringEditRanges) {

}

func (
	stringEditor *StringEditors) Substitute(
	editRange StringEditRanges,
	stringToSubstitute string) {

}

func (
	stringEditor *StringEditors) GetCurrentString() CurrentStrings {

	var currentString CurrentStrings

	return currentString
}

func (
	stringEditor *StringEditors) GetStringEditHistory() *StringEditHistories {

	return stringEditor.StringEditHistory
}
