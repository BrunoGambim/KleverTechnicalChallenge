package comment_repository

import models "KleverTechnicalChallenge/domain/models"

type CommentRepository interface {
	FindAll() ([]models.Comment, error)
	Insert(comment models.Comment) (string, error)
	Lock()
	Unlock()
}
