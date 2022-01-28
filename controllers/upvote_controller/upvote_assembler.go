package upvote_controller

import (
	"KleverTechnicalChallenge/domain/models"

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
	return &GetUpvoteDTO{
		Id:        upvote.Id.Hex(),
		Type:      upvote.Type,
		CommentId: upvote.Id.Hex(),
	}
}
