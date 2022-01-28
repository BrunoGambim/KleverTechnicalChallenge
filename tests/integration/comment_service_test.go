package services_test

import (
	comment_repository "KleverTechnicalChallenge/database/repositories/comment_repository"
	"KleverTechnicalChallenge/domain/models"
	"KleverTechnicalChallenge/domain/services"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var commentService *services.CommentService

func init() {
	godotenv.Overload(".env")
	repository, _ := comment_repository.NewCommentRepository()
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
