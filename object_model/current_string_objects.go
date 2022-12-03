package object_model

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

//#TODO remove Characters from CurrentString (just links)
//#TODO should we just have the start link pointer
//#TODO go for a pointer approach

type CurrentStrings struct {
	objects.BnogObjects
	Characters             []CharacterTokens
	CurrentStringLinksList []CurrentStringLinks
}

func (currentString *CurrentStrings) Initialise(
	stringToBeCreated string) {

	stringLength :=
		len(
			stringToBeCreated)

	currentString.Characters =
		make(
			[]CharacterTokens,
			stringLength)

	currentString.CurrentStringLinksList =
		make(
			[]CurrentStringLinks,
			stringLength-1)

	appendCharactersAndLinks(
		stringToBeCreated,
		currentString)

}

func appendCharactersAndLinks(
	stringToBeCreated string,
	currentStrings *CurrentStrings) {

	for characterPosition, character := range stringToBeCreated {

		appendCharacter(
			currentStrings,
			characterPosition,
			character)

		append_link(
			currentStrings,
			characterPosition)
	}
}

func appendCharacter(
	currentString *CurrentStrings,
	characterPosition int,
	character int32) {

	currentString.Characters[characterPosition].character_uuid, _ =
		uuid_service.GetUUID(
			1,
			"")

	currentString.
		Characters[characterPosition].
		Code_point =
		character
}

func append_link(
	current_string *CurrentStrings,
	character_position int) {

	if character_position == 0 {
		return
	}

	current_string_link :=
		&current_string.
			CurrentStringLinksList[character_position-1]

	current_string_link.
		Previous_character_token =
		current_string.
			Characters[character_position-1]

	current_string_link.
		Next_character_token =
		current_string.
			Characters[character_position]

}

func (currentString *CurrentStrings) GetCurrentStringArray() []rune {

	current_string_length := len(currentString.Characters)

	current_string_array := make(
		[]rune,
		current_string_length)

	for index, character := range currentString.Characters {
		current_string_array[index] = character.Code_point
	}

	return current_string_array

}
