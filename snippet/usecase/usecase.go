package usecase

import (
	"github.com/jordyf15/code-sharing-app/custom_errors"
	"github.com/jordyf15/code-sharing-app/models"
	"github.com/jordyf15/code-sharing-app/snippet"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type snippetUsecase struct {
	snippetRepo snippet.Repository
}

func NewSnippetUsecase(snippetRepo snippet.Repository) snippet.Usecase {
	return &snippetUsecase{snippetRepo: snippetRepo}
}

func (usecase *snippetUsecase) Create(code, language, theme string) (*models.Snippet, error) {
	snippet := &models.Snippet{}
	snippet.Code = code
	snippet.SetLanguage(language)
	snippet.SetTheme(theme)
	errors := snippet.VerifyFields()
	if len(errors) > 0 {
		return nil, &custom_errors.MultipleErrors{Errors: errors}
	}

	err := usecase.snippetRepo.Create(snippet)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}

func (usecase *snippetUsecase) GetSnippet(snippetID primitive.ObjectID) (*models.Snippet, error) {
	return usecase.snippetRepo.GetByID(snippetID)
}

func (usecase *snippetUsecase) UpdateSnippet(snippetID primitive.ObjectID, updates map[string]interface{}) (*models.Snippet, error) {
	snippet, err := usecase.snippetRepo.GetByID(snippetID)
	if err != nil {
		return nil, err
	}

	if updates["language"] != nil {
		snippet.SetLanguage(updates["language"].(string))
	}
	if updates["theme"] != nil {
		snippet.SetTheme(updates["theme"].(string))
	}
	if updates["code"] != nil {
		snippet.Code = updates["code"].(string)
	}
	errors := snippet.VerifyFields()
	if len(errors) > 0 {
		return nil, &custom_errors.MultipleErrors{Errors: errors}
	}

	err = usecase.snippetRepo.Update(snippetID, updates)
	if err != nil {
		return nil, err
	}

	return snippet, nil
}
