package services

import (
	repositories "KleverTechnicalChallenge/database/repositories"
	models "KleverTechnicalChallenge/domain/models"
)

type UpvoteService struct {
	upvoteRepository *repositories.UpvoteRepository
}

func NewUpvoteService() (*UpvoteService, error) {
	repository, err := repositories.NewUpvoteRepository()
	if err != nil {
		return &UpvoteService{}, err
	}
	return &UpvoteService{
		upvoteRepository: repository,
	}, nil
}

func (service *UpvoteService) FindById(id string) (models.Upvote, error) {
	service.upvoteRepository.Lock()
	defer service.upvoteRepository.Unlock()
	upvote, err := service.upvoteRepository.FindById(id)
	return upvote, err
}

func (service *UpvoteService) Insert(upvote models.Upvote) (string, error) {
	service.upvoteRepository.Lock()
	defer service.upvoteRepository.Unlock()
	id, err := service.upvoteRepository.Insert(upvote)
	return id, err
}

func (service *UpvoteService) DeleteById(id string) error {
	service.upvoteRepository.Lock()
	defer service.upvoteRepository.Unlock()
	err := service.upvoteRepository.DeleteById(id)
	return err
}
