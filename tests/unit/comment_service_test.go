package services_unit_test

import (
	"KleverTechnicalChallenge/domain/models"
	"KleverTechnicalChallenge/domain/services"

	"testing"

	mocked_repositories "KleverTechnicalChallenge/tests/mocked_repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var commentService *services.CommentService

func init() {
	repository := mocked_repositories.NewCommentRepository()
	commentService, _ = services.NewCommentService(repository)
}

func TestCommentService(t *testing.T) {
	comments, _ := commentService.FindAll()
	numberOfDocumentsBeforeInsertion := len(comments)
	id, _ := commentService.Insert(models.Comment{
		Message: "comentario de teste",
	})
	comments, _ = commentService.FindAll()
	numberOfDocumentsAfterInsertion := len(comments)

	if numberOfDocumentsBeforeInsertion+1 != numberOfDocumentsAfterInsertion {
		t.Errorf("Erro ao inserir ou ler dados na base de dados")
	}

	if !primitive.IsValidObjectID(id) {
		t.Errorf("O id retornado pela insercao é invalido")
	}
}
