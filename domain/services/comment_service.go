package services

import (
	comment_repository "KleverTechnicalChallenge/database/repositories/comment_repository"
	models "KleverTechnicalChallenge/domain/models"
)

type CommentService struct {
	commentRepository comment_repository.CommentRepository
}

func NewCommentService() (*CommentService, error) {
	repository, err := comment_repository.NewCommentRepository()
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
