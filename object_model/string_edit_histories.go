package object_model

type StringEditHistories struct {
	Four_d_string_links   *[]FourDStringLinks
	string_history        []string
	string_marked_history []string
	original_string       string
	current_string        string
	modified_string       string
	marked_string         string
}

func (string_edit_history *StringEditHistories) Create(initial_string string) {

	string_edit_history.
		original_string = initial_string

	string_edit_history.
		current_string = initial_string

}

func (string_edit_history *StringEditHistories) Terminate() {

}

func (string_edit_history *StringEditHistories) GetEditTransactions() []int {
	var edit_transactions []int

	return edit_transactions
}

func (string_edit_history *StringEditHistories) GetString(transaction_id int) string {

	if transaction_id == 0 {
		return string_edit_history.original_string
	}

	return ""
}

func (string_edit_history *StringEditHistories) GetCurrentString() string {

	return string_edit_history.current_string

}

func (string_edit_history *StringEditHistories) SetCurrentString(string_to_set string) {

	string_edit_history.string_history =
		append(
			string_edit_history.string_history,
			string_edit_history.current_string)

	string_edit_history.current_string = string_to_set
}

func (string_edit_history *StringEditHistories) GetStringEditHistory() []string {

	return string_edit_history.string_history
}

func (string_edit_history *StringEditHistories) GetEditChanges(transaction_id int) string {

	var marked_string string

	return marked_string
}

func (string_edit_history *StringEditHistories) Get_modified_string() string {

	return string_edit_history.modified_string
}

func (string_edit_history *StringEditHistories) Get_marked_string() string {

	return string_edit_history.marked_string
}

//TODO - Temporary solution for setting history
func (string_edit_history *StringEditHistories) Set_string_changes(modified_string, marked_string string) {

	string_edit_history.modified_string = modified_string
	string_edit_history.marked_string = marked_string
	string_edit_history.string_marked_history = append(string_edit_history.string_marked_history, marked_string)

}
