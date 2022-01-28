package comment_repository

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

var commentRepositoryInstance *CommentRepositoryImpl
var commentRepositoryInstanceError error
var commentRepositoryOnce sync.Once

type CommentRepositoryImpl struct {
	sync.Mutex
	collection *mongo.Collection
	ctx        context.Context
}

func NewCommentRepository() (*CommentRepositoryImpl, error) {
	commentRepositoryOnce.Do(func() {
		ctx := context.Background()
		client, err := connectionFactory.GetMongoClient(ctx)

		if err != nil {
			commentRepositoryInstance = &CommentRepositoryImpl{}
			commentRepositoryInstanceError = err
		}

		databaseName := os.Getenv("DATABASE_NAME")
		commentCollection := os.Getenv("COMMENTS_COLLECTION")
		commentRepositoryInstance = &CommentRepositoryImpl{
			collection: client.Database(databaseName).Collection(commentCollection),
			ctx:        ctx,
		}
		commentRepositoryInstanceError = nil
	})
	return commentRepositoryInstance, commentRepositoryInstanceError
}

func (repository *CommentRepositoryImpl) Insert(comment models.Comment) (string, error) {
	result, err := repository.collection.InsertOne(repository.ctx, comment)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *CommentRepositoryImpl) FindAll() ([]models.Comment, error) {
	filter := bson.M{}

	cur, err := repository.collection.Find(repository.ctx, filter)
	if err != nil {
		return []models.Comment{}, err
	}

	result, err := repository.commentListFromCur(cur)

	return result, err
}

func (repository *CommentRepositoryImpl) commentListFromCur(cur *mongo.Cursor) ([]models.Comment, error) {
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
