package object_model

import (
	"core_foundation/core_object_model"
	"database_manager/utils"
)

//#TODO remove Characters from CurrentString (just links)
//#TODO should we just have the start link pointer
//#TODO go for a pointer approach

type CurrentStrings struct {
	core_object_model.Objects
	Characters                []CharacterTokens
	Current_string_links_list []CurrentStringLinks
}

func (current_string *CurrentStrings) Initialise(
	string_to_be_created string) {

	string_length :=
		len(
			string_to_be_created)

	current_string.Characters =
		make(
			[]CharacterTokens,
			string_length)

	current_string.Current_string_links_list =
		make(
			[]CurrentStringLinks,
			string_length-1)

	append_characters_and_links(
		string_to_be_created,
		current_string)

}

func append_characters_and_links(
	string_to_be_created string,
	current_string *CurrentStrings) {

	for character_position, character := range string_to_be_created {

		append_character(
			current_string,
			character_position,
			character)

		append_link(
			current_string,
			character_position)
	}
}

func append_character(
	current_string *CurrentStrings,
	character_position int,
	character int32) {

	current_string.Characters[character_position].character_uuid, _ =
		utils.GetUUID(
			1,
			"")

	current_string.
		Characters[character_position].
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
			Current_string_links_list[character_position-1]

	current_string_link.
		Previous_character_token =
		current_string.
			Characters[character_position-1]

	current_string_link.
		Next_character_token =
		current_string.
			Characters[character_position]

}

func (current_string *CurrentStrings) GetCurrentStringArray() []rune {

	current_string_length := len(current_string.Characters)

	current_string_array := make(
		[]rune,
		current_string_length)

	for index, character := range current_string.Characters {
		current_string_array[index] = character.Code_point
	}

	return current_string_array

}
