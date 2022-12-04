package testing

import (
	"fmt"
	"github.com/OntoLedgy/string_editing_services/object_model"
	"testing"
)

func TestCurrentString(t *testing.T) {

	var currentString object_model.CurrentStrings

	currentString.Initialise("test")

	fmt.Printf(
		"%v\n%v\n",
		currentString.GetCurrentStringArray(),
		currentString.CurrentStringLinksList)

}
