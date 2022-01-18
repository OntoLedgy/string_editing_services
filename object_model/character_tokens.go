package object_model

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type CharacterTokens struct {
	object         *objects.BnogObjects
	character_uuid *utils.UUIDs
	Code_point     rune
}
