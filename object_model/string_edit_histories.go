package object_model

type StringEditHistories struct {
	FourDimensionalStringLinks *[]FourDStringLinks
	stringHistory              []string
	stringMarkedHistory        []string
	originalString             string
	currentString              string
	modifiedString             string
	markedString               string
}

func (stringEditHistory *StringEditHistories) Create(initialString string) {

	stringEditHistory.
		originalString = initialString

	stringEditHistory.
		currentString = initialString

}

func (stringEditHistory *StringEditHistories) Terminate() {

}

func (stringEditHistory *StringEditHistories) GetEditTransactions() []int {
	var editTransactions []int

	return editTransactions
}

func (stringEditHistory *StringEditHistories) GetString(transactionId int) string {

	if transactionId == 0 {
		return stringEditHistory.originalString
	}

	return ""
}

func (stringEditHistory *StringEditHistories) GetCurrentString() string {

	return stringEditHistory.currentString

}

func (stringEditHistory *StringEditHistories) SetCurrentString(stringToSet string) {

	stringEditHistory.stringHistory =
		append(
			stringEditHistory.stringHistory,
			stringEditHistory.currentString)

	stringEditHistory.currentString = stringToSet
}

func (stringEditHistory *StringEditHistories) GetStringEditHistory() []string {

	return stringEditHistory.stringHistory
}

func (stringEditHistory *StringEditHistories) GetEditChanges(transactionId int) string {

	var markedString string

	return markedString
}

func (stringEditHistory *StringEditHistories) GetModifiedString() string {

	return stringEditHistory.modifiedString
}

func (stringEditHistory *StringEditHistories) GetMarkedString() string {

	return stringEditHistory.markedString
}

//TODO - Temporary solution for setting history
func (stringEditHistory *StringEditHistories) SetStringChanges(modifiedString, markedString string) {

	stringEditHistory.modifiedString = modifiedString
	stringEditHistory.markedString = markedString
	stringEditHistory.stringMarkedHistory = append(stringEditHistory.stringMarkedHistory, markedString)

}
