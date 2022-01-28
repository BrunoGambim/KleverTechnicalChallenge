package repositories

import (
	connectionFactory "KleverTechnicalChallenge/database/connection"
	models "KleverTechnicalChallenge/domain/models"
	"os"

	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var commentRepositoryInstance *CommentRepository
var commentRepositoryInstanceError error
var commentRepositoryOnce sync.Once

type CommentRepository struct {
	sync.Mutex
	collection *mongo.Collection
	ctx        context.Context
}

func NewCommentRepository() (*CommentRepository, error) {
	commentRepositoryOnce.Do(func() {
		ctx := context.Background()
		client, err := connectionFactory.GetMongoClient(ctx)

		if err != nil {
			commentRepositoryInstance = &CommentRepository{}
			commentRepositoryInstanceError = err
		}

		databaseName := os.Getenv("DATABASE_NAME")
		commentsCollection := os.Getenv("COMMENTS_COLLECTION")

		commentRepositoryInstance = &CommentRepository{
			collection: client.Database(databaseName).Collection(commentsCollection),
			ctx:        ctx,
		}
		commentRepositoryInstanceError = nil
	})
	return commentRepositoryInstance, commentRepositoryInstanceError
}

func (repository *CommentRepository) Insert(comment models.Comment) (string, error) {
	result, err := repository.collection.InsertOne(repository.ctx, comment)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *CommentRepository) FindAll() ([]models.Comment, error) {
	filter := bson.M{}

	cur, err := repository.collection.Find(repository.ctx, filter)
	if err != nil {
		return []models.Comment{}, err
	}

	result, err := repository.commentListFromCur(cur)

	return result, err
}

func (repository *CommentRepository) commentListFromCur(cur *mongo.Cursor) ([]models.Comment, error) {
	result := []models.Comment{}
	for cur.Next(repository.ctx) {
		comment := models.Comment{}
		err := cur.Decode(&comment)
		if err != nil {
			return result, err
		}
		result = append(result, comment)
	}
	return result, nil
}
