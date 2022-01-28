package repositories

import (
	connectionFactory "KleverTechnicalChallenge/database/connection"
	"KleverTechnicalChallenge/domain/models"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UPVOTES_COLLECTION = "upvotes"
)

var upvoteRepositoryInstance *UpvoteRepository
var upvoteRepositoryInstanceError error
var upvoteRepositoryOnce sync.Once

type UpvoteRepository struct {
	sync.Mutex
	collection *mongo.Collection
	ctx        context.Context
}

func NewUpvoteRepository() (*UpvoteRepository, error) {
	upvoteRepositoryOnce.Do(func() {
		ctx := context.Background()
		client, err := connectionFactory.GetMongoClient(ctx)

		if err != nil {
			upvoteRepositoryInstance = &UpvoteRepository{}
			upvoteRepositoryInstanceError = err
		}

		upvoteRepositoryInstance = &UpvoteRepository{
			collection: client.Database(DATABASE_NAME).Collection(UPVOTES_COLLECTION),
			ctx:        ctx,
		}
		upvoteRepositoryInstanceError = nil
	})
	return upvoteRepositoryInstance, upvoteRepositoryInstanceError
}

func (repository *UpvoteRepository) FindById(id string) (models.Upvote, error) {
	result := models.Upvote{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": objectId}
	err = repository.collection.FindOne(repository.ctx, filter).Decode(&result)
	return result, err
}

func (repository *UpvoteRepository) Insert(upvote models.Upvote) (string, error) {
	result, err := repository.collection.InsertOne(repository.ctx, upvote)
	if err != nil {
		return "", err
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, err
}

func (repository *UpvoteRepository) DeleteById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	_, err = repository.collection.DeleteOne(repository.ctx, filter)
	return err
}
