package testing

import (
	"fmt"
	"github.com/OntoLedgy/string_editing_services/factories"
	"testing"
)

func TestFactories(t *testing.T) {

	var stringEditorFactory factories.StringEditorFactory

	stringEditor := stringEditorFactory.CreateStringEditor(
		"test1")

	edit_range := factories.CreateStringEditRange(1, 2)

	fmt.Printf("%v\n%v\n", stringEditor, edit_range)
}
