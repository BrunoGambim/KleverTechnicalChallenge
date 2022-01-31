package upvote_repository

import (
	"KleverTechnicalChallenge/domain/models"
)

type UpvoteRepository interface {
	FindById(id string) (models.Upvote, error)
	FindByCommentId(commentId string) ([]models.Upvote, error)
	Insert(upvote models.Upvote) (string, error)
	DeleteById(id string) error
}
