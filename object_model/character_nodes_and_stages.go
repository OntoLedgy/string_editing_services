package object_model

type CharacterNodesAndStages struct {
	representedCharacterToken *CharacterTokens
}

type CharacterNodes struct {
	*CharacterNodesAndStages
}

type CharacterStages struct {
	*CharacterNodesAndStages
	parentStage *CharacterNodesAndStages
}

type CharacterBoundaries struct{}

type FourDStringLinks struct {
	PreviousCharacterNode CharacterTokens
	NextCharacterNode     CharacterTokens
}

type BeforeAfterLinks struct{}
