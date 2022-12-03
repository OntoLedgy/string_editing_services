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

		appendLink(
			currentStrings,
			characterPosition)
	}
}

func appendCharacter(
	currentString *CurrentStrings,
	characterPosition int,
	character int32) {

	currentString.Characters[characterPosition].characterUuid, _ =
		uuid_service.GetUUID(
			1,
			"")

	currentString.
		Characters[characterPosition].
		CodePoint =
		character
}

func appendLink(
	currentString *CurrentStrings,
	characterPosition int) {

	if characterPosition == 0 {
		return
	}

	currentStringLink :=
		&currentString.
			CurrentStringLinksList[characterPosition-1]

	currentStringLink.
		PreviousCharacterToken =
		currentString.
			Characters[characterPosition-1]

	currentStringLink.
		NextCharacterToken =
		currentString.
			Characters[characterPosition]

}

func (currentString *CurrentStrings) GetCurrentStringArray() []rune {

	currentStringLength := len(currentString.Characters)

	currentStringArray := make(
		[]rune,
		currentStringLength)

	for index, character := range currentString.Characters {
		currentStringArray[index] = character.CodePoint
	}

	return currentStringArray

}
