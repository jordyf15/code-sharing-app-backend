package repository

import (
	"context"
	"time"

	"github.com/jordyf15/code-sharing-app/models"
	"github.com/jordyf15/code-sharing-app/snippet"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const contextTimeout = time.Second * 30

type snippetRepository struct {
	db *mongo.Collection
}

func NewSnippetRepository(db *mongo.Database) snippet.Repository {
	collection := db.Collection("snippets")
	return &snippetRepository{db: collection}
}

func (repo *snippetRepository) GetByID(snippetID primitive.ObjectID) (*models.Snippet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	filter := bson.D{
		{Key: "_id", Value: snippetID},
	}

	foundSnippet := &models.Snippet{}
	err := repo.db.FindOne(ctx, filter).Decode(foundSnippet)

	return foundSnippet, err
}

func (repo *snippetRepository) Update(snippetID primitive.ObjectID, updates map[string]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)

	defer cancel()

	updatesValue := bson.D{{Key: "updated_at", Value: time.Now()}}
	for k, v := range updates {
		updatesValue = append(updatesValue, bson.E{Key: k, Value: v})
	}

	filter := bson.D{{Key: "_id", Value: snippetID}}
	updatesBSON := bson.D{{Key: "$set", Value: updatesValue}}
	updateResult, err := repo.db.UpdateOne(ctx, filter, updatesBSON)

	if updateResult.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return err
}

func (repo *snippetRepository) Create(snippet *models.Snippet) error {
	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	snippet.ID = primitive.NewObjectID()
	snippet.CreatedAt = time.Now()
	snippet.UpdatedAt = time.Now()
	_, err := repo.db.InsertOne(ctx, snippet)

	return err
}
