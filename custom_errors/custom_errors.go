package custom_errors

import "strings"

var (
	// general errors
	ErrUnknownErrorOccured = newErr(101, "Unknown error occured")
	ErrRecordNotFound      = newErr(102, "Record not found")

	// Snippet Errors
	ErrSnippetLanguageInvalid = newErr(201, "Invalid snippet language")
	ErrSnippetThemeInvalid    = newErr(202, "Invalid snippet theme")
)

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (err *Error) Error() string {
	return err.Message
}

func newErr(code int, message string) *Error {
	return &Error{Message: message, Code: code}
}

type MultipleErrors struct {
	Errors []error `json:"errors"`
}

func (multipleErr *MultipleErrors) Error() string {
	messages := make([]string, len(multipleErr.Errors))
	for i, error := range multipleErr.Errors {
		messages[i] = error.Error()
	}

	return strings.Join(messages, ", ")
}
