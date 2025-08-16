package snippet

import (
	"github.com/jordyf15/code-sharing-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(snippet *models.Snippet) error
	Update(snippetID primitive.ObjectID, updates map[string]interface{}) error
	GetByID(snippetID primitive.ObjectID) (*models.Snippet, error)
}

type Usecase interface {
	Create(code, language, theme string) (*models.Snippet, error)
	GetSnippet(snippetID primitive.ObjectID) (*models.Snippet, error)
	UpdateSnippet(snippetID primitive.ObjectID, updates map[string]interface{}) (*models.Snippet, error)
}
