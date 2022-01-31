package mocked_repositories

import (
	"KleverTechnicalChallenge/domain/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockedUpvoteRepository struct {
	collection []models.Upvote
}

func NewUpvoteRepository() *MockedUpvoteRepository {
	return &MockedUpvoteRepository{
		collection: []models.Upvote{},
	}
}

func (repository *MockedUpvoteRepository) FindById(id string) (models.Upvote, error) {
	for _, upvote := range repository.collection {
		if upvote.Id.Hex() == id {
			return upvote, nil
		}
	}
	return models.Upvote{}, mongo.ErrNoDocuments
}

func (repository *MockedUpvoteRepository) FindByCommentId(commentId string) ([]models.Upvote, error) {
	result := []models.Upvote{}
	for _, upvote := range repository.collection {
		if upvote.CommentId.Hex() == commentId {
			result = append(result, upvote)
		}
	}
	return result, nil
}

func (repository *MockedUpvoteRepository) Insert(upvote models.Upvote) (string, error) {
	newObjId := primitive.NewObjectIDFromTimestamp(time.Now())
	upvote.Id = newObjId
	repository.collection = append(repository.collection, upvote)
	return newObjId.Hex(), nil
}

func (repository *MockedUpvoteRepository) DeleteById(id string) error {
	for i, upvote := range repository.collection {
		if upvote.Id.Hex() == id {
			repository.collection[i] = repository.collection[len(repository.collection)-1]
			repository.collection = repository.collection[:len(repository.collection)-1]
		}
	}
	return nil
}
