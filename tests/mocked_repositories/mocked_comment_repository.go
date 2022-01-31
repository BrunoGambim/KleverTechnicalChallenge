package mocked_repositories

import (
	models "KleverTechnicalChallenge/domain/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockedCommentRepository struct {
	collection []models.Comment
}

func NewCommentRepository() *MockedCommentRepository {
	return &MockedCommentRepository{
		collection: []models.Comment{},
	}
}

func (repository *MockedCommentRepository) Insert(comment models.Comment) (string, error) {
	newObjId := primitive.NewObjectIDFromTimestamp(time.Now())
	comment.Id = newObjId
	repository.collection = append(repository.collection, comment)
	return newObjId.Hex(), nil
}

func (repository *MockedCommentRepository) FindAll() ([]models.Comment, error) {
	return repository.collection, nil
}
