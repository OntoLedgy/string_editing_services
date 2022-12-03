package object_model

type CharacterNodesAndStages struct {
	representedCharacterToken *CharacterTokens
}

type CharacterNodes struct {
	*CharacterNodesAndStages
}

type CharacterStages struct {
	*CharacterNodesAndStages
	parent_stage *CharacterNodesAndStages
}

type CharacterBoundaries struct{}

type FourDStringLinks struct {
	Previous_character_node CharacterTokens
	Next_character_node     CharacterTokens
}

type BeforeAfterLinks struct{}
