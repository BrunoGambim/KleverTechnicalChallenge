package services

import (
	repositories "KleverTechnicalChallenge/database/repositories"
	models "KleverTechnicalChallenge/domain/models"
)

type CommentService struct {
	commentRepository *repositories.CommentRepository
}

func NewCommentService() (*CommentService, error) {
	repository, err := repositories.NewCommentRepository()
	if err != nil {
		return &CommentService{}, err
	}
	return &CommentService{
		commentRepository: repository,
	}, nil
}

func (service *CommentService) FindAll() ([]models.Comment, error) {
	service.commentRepository.Lock()
	defer service.commentRepository.Unlock()
	comments, err := service.commentRepository.FindAll()
	return comments, err
}

func (service *CommentService) Insert(comment models.Comment) (string, error) {
	service.commentRepository.Lock()
	defer service.commentRepository.Unlock()
	id, err := service.commentRepository.Insert(comment)
	return id, err
}
