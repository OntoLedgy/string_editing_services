package object_model

import (
	"github.com/OntoLedgy/core_ontology/core_object_model"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type CharacterTokens struct {
	object         *core_object_model.Objects
	character_uuid *utils.UUIDs
	Code_point     rune
}
