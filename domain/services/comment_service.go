package services

import (
	comment_repository "KleverTechnicalChallenge/database/repositories/comment_repository"
	models "KleverTechnicalChallenge/domain/models"
)

type CommentService struct {
	commentRepository comment_repository.CommentRepository
}

func NewCommentService(repository comment_repository.CommentRepository) (*CommentService, error) {
	return &CommentService{
		commentRepository: repository,
	}, nil
}

func (service *CommentService) FindAll() ([]models.Comment, error) {
	comments, err := service.commentRepository.FindAll()
	return comments, err
}

func (service *CommentService) Insert(comment models.Comment) (string, error) {
	id, err := service.commentRepository.Insert(comment)
	return id, err
}
