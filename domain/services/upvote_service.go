package services

import (
	upvote_repository "KleverTechnicalChallenge/database/repositories/upvote_repository"
	models "KleverTechnicalChallenge/domain/models"
	"time"
)

type UpvoteService struct {
	upvoteRepository upvote_repository.UpvoteRepository
}

func NewUpvoteService(repository upvote_repository.UpvoteRepository) (*UpvoteService, error) {
	return &UpvoteService{
		upvoteRepository: repository,
	}, nil
}

func (service *UpvoteService) FindById(id string) ([]models.Upvote, error) {
	upvote, err := service.upvoteRepository.FindById(id)
	return upvote, err
}

func (service *UpvoteService) FindByCommentId(commentId string) ([]models.Upvote, error) {
	upvote, err := service.upvoteRepository.FindByCommentId(commentId)
	return upvote, err
}

func (service *UpvoteService) Insert(upvote models.Upvote) (string, error) {
	upvote.CreatedAt = uint64(time.Now().Unix())
	id, err := service.upvoteRepository.Insert(upvote)
	return id, err
}

func (service *UpvoteService) DeleteById(id string) error {
	err := service.upvoteRepository.DeleteById(id)
	return err
}
