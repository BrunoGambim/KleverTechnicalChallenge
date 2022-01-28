package upvote_controller

import (
	"KleverTechnicalChallenge/domain/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func upvoteFromCreateDto(upvoteDto *CreateUpvoteDTO) (models.Upvote, error) {
	commentId, err := primitive.ObjectIDFromHex(upvoteDto.CommentId)
	return models.Upvote{
		Type:      upvoteDto.Type.Enum().String(),
		CommentId: commentId,
	}, err
}

func getAlbumDtoFromModel(upvote models.Upvote) *GetUpvoteDTO {
	createdAtTime := time.Unix(int64(upvote.CreatedAt), 0)
	createdAt := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		createdAtTime.Year(), createdAtTime.Month(), createdAtTime.Day(),
		createdAtTime.Hour(), createdAtTime.Minute(), createdAtTime.Second())
	return &GetUpvoteDTO{
		Id:        upvote.Id.Hex(),
		Type:      upvote.Type,
		CommentId: upvote.Id.Hex(),
		CreatedAt: createdAt,
	}
}
