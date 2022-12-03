package testing

import (
	"fmt"
	"github.com/OntoLedgy/string_editing_services/object_model"
	"testing"
)

func TestCurrentString(t *testing.T) {

	var current_string object_model.CurrentStrings

	current_string.Initialise("test")

	fmt.Printf(
		"%v\n%v\n",
		current_string.GetCurrentStringArray(),
		current_string.CurrentStringLinksList)

}
