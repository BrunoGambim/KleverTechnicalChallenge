package upvote_repository

import (
	connectionFactory "KleverTechnicalChallenge/database/connection"
	"KleverTechnicalChallenge/domain/models"
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var upvoteRepositoryInstance *UpvoteRepositoryImpl
var upvoteRepositoryInstanceError error
var upvoteRepositoryOnce sync.Once

type UpvoteRepositoryImpl struct {
	sync.Mutex
	collection *mongo.Collection
	ctx        context.Context
}

func NewUpvoteRepository() (*UpvoteRepositoryImpl, error) {
	upvoteRepositoryOnce.Do(func() {
		ctx := context.Background()
		client, err := connectionFactory.GetMongoClient(ctx)

		if err != nil {
			upvoteRepositoryInstance = &UpvoteRepositoryImpl{}
			upvoteRepositoryInstanceError = err
		}

		databaseName := os.Getenv("DATABASE_NAME")
		upvotesCollection := os.Getenv("UPVOTES_COLLECTION")
		upvoteRepositoryInstance = &UpvoteRepositoryImpl{
			collection: client.Database(databaseName).Collection(upvotesCollection),
			ctx:        ctx,
		}
		upvoteRepositoryInstanceError = nil
	})
	return upvoteRepositoryInstance, upvoteRepositoryInstanceError
}

func (repository *UpvoteRepositoryImpl) FindById(id string) ([]models.Upvote, error) {
	repository.Lock()
	defer repository.Unlock()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return []models.Upvote{}, err
	}

	filter := bson.M{"_id": objectId}
	cur, err := repository.collection.Find(repository.ctx, filter)

	if err != nil {
		return []models.Upvote{}, err
	}

	result, err := repository.upvoteListFromCur(cur)
	return result, err
}

func (repository *UpvoteRepositoryImpl) FindByCommentId(commentId string) ([]models.Upvote, error) {
	repository.Lock()
	defer repository.Unlock()

	objectId, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return []models.Upvote{}, err
	}

	filter := bson.M{"comment_id": objectId}
	cur, err := repository.collection.Find(repository.ctx, filter)

	if err != nil {
		return []models.Upvote{}, err
	}

	result, err := repository.upvoteListFromCur(cur)
	return result, err
}

func (repository *UpvoteRepositoryImpl) upvoteListFromCur(cur *mongo.Cursor) ([]models.Upvote, error) {
	result := []models.Upvote{}
	for cur.Next(repository.ctx) {
		upvote := models.Upvote{}
		err := cur.Decode(&upvote)
		if err != nil {
			return result, err
		}
		result = append(result, upvote)
	}
	return result, nil
}

func (repository *UpvoteRepositoryImpl) Insert(upvote models.Upvote) (string, error) {
	repository.Lock()
	defer repository.Unlock()

	result, err := repository.collection.InsertOne(repository.ctx, upvote)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *UpvoteRepositoryImpl) DeleteById(id string) error {
	repository.Lock()
	defer repository.Unlock()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	_, err = repository.collection.DeleteOne(repository.ctx, filter)
	return err
}
