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

func (repository *UpvoteRepository) FindById(id string) ([]models.Upvote, error) {
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

func (repository *UpvoteRepository) upvoteListFromCur(cur *mongo.Cursor) ([]models.Upvote, error) {
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
