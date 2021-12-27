package object_model

import (
	"core_foundation/core_object_model"
	"database_manager/utils"
)

type CharacterTokens struct {
	object         *core_object_model.Objects
	character_uuid *utils.UUIDs
	Code_point     rune
}
