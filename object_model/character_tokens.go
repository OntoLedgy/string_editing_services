package object_model

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type CharacterTokens struct {
	object         *objects.BnogObjects
	character_uuid *uuid_service.UUIDs
	Code_point     rune
}
