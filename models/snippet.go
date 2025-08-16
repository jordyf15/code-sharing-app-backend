package models

import (
	"encoding/json"
	"time"

	"github.com/jordyf15/code-sharing-app/custom_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SnippetLanguage string

const (
	SnippetLanguageHTML    = "HTML"
	SnippetLanguageJS      = "Javascript"
	SnippetLanguageCSS     = "CSS"
	SnippetLanguageInvalid = "invalid"
)

type SnippetTheme string

const (
	SnippetThemeLight   = "Light"
	SnippetThemeDark    = "Dark"
	SnippetThemeInvalid = "invalid"
)

type Snippet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Code      string             `json:"code" bson:"code"`
	Language  SnippetLanguage    `json:"language" bson:"language"`
	Theme     SnippetTheme       `json:"theme" bson:"theme"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func (snippet *Snippet) VerifyFields() []error {
	errors := make([]error, 0)

	if snippet.Theme != SnippetThemeLight && snippet.Theme != SnippetThemeDark {
		errors = append(errors, custom_errors.ErrSnippetThemeInvalid)
	}
	if snippet.Language != SnippetLanguageHTML && snippet.Language != SnippetLanguageJS && snippet.Language != SnippetLanguageCSS {
		errors = append(errors, custom_errors.ErrSnippetLanguageInvalid)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (snippet *Snippet) SetLanguage(language string) {
	switch language {
	case SnippetLanguageHTML:
		snippet.Language = SnippetLanguageHTML
	case SnippetLanguageCSS:
		snippet.Language = SnippetLanguageCSS
	case SnippetLanguageJS:
		snippet.Language = SnippetLanguageJS
	default:
		snippet.Language = SnippetLanguageInvalid
	}
}

func (snippet *Snippet) SetTheme(theme string) {
	switch theme {
	case SnippetThemeLight:
		snippet.Theme = SnippetThemeLight
	case SnippetThemeDark:
		snippet.Theme = SnippetThemeDark
	default:
		snippet.Theme = SnippetThemeInvalid
	}
}

func (snippet *Snippet) MarshalJSON() ([]byte, error) {
	type Alias Snippet
	newStruct := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		Alias: (*Alias)(snippet),
	}

	newStruct.CreatedAt = snippet.CreatedAt.Format("2006-01-02T15:04:05-0700")
	newStruct.UpdatedAt = snippet.UpdatedAt.Format("2006-01-02T15:04:05-0700")

	return json.Marshal(newStruct)
}

func (snippet *Snippet) UnmarshalJSON(data []byte) error {
	type Alias Snippet
	alias := &struct {
		*Alias
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{Alias: (*Alias)(snippet)}

	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	snippet.CreatedAt, err = time.Parse("2006-01-02T15:04:05-0700", alias.CreatedAt)
	if err != nil {
		return err
	}

	snippet.UpdatedAt, err = time.Parse("2006-01-02T15:04:05-0700", alias.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
