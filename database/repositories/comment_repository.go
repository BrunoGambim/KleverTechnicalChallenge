package repositories

import (
	connectionFactory "KleverTechnicalChallenge/database/connection"
	models "KleverTechnicalChallenge/domain/models"

	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE_NAME       = "KleverTechnicalChallenge"
	COMMENTS_COLLECTION = "comments"
)

var repositoryInstance *CommentRepository
var repositoryInstanceError error
var repositoryOnce sync.Once

type CommentRepository struct {
	sync.Mutex
	collection *mongo.Collection
	context    context.Context
}

func NewCommentRepository() (*CommentRepository, error) {
	repositoryOnce.Do(func() {
		context := context.Background()
		client, err := connectionFactory.GetMongoClient(context)

		if err != nil {
			repositoryInstance = &CommentRepository{}
			repositoryInstanceError = err
		}

		repositoryInstance = &CommentRepository{
			collection: client.Database(DATABASE_NAME).Collection(COMMENTS_COLLECTION),
			context:    context,
		}
		repositoryInstanceError = nil
	})
	return repositoryInstance, repositoryInstanceError
}

func (repository *CommentRepository) Insert(comment models.Comment) (string, error) {
	result, err := repository.collection.InsertOne(repository.context, comment)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *CommentRepository) FindAll() ([]models.Comment, error) {
	comments := []models.Comment{}
	filter := bson.M{}

	cur, err := repository.collection.Find(repository.context, filter)
	if err != nil {
		return comments, err
	}

	for cur.Next(repository.context) {
		comment := models.Comment{}
		err := cur.Decode(&comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, err
}
