package testing

import (
	"fmt"

	"github.com/OntoLedgy/string_editing_services/factories"
	"testing"
)

func TestFactories(t *testing.T) {

	var string_editor_factory factories.StringEditorFactory

	string_editor := string_editor_factory.CreateStringEditor(
		"test1")

	edit_range := factories.CreateStringEditRange(1, 2)

	fmt.Printf("%v\n%v\n", string_editor, edit_range)
}
