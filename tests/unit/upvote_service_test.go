package services_unit_test

import (
	"KleverTechnicalChallenge/domain/models"
	"KleverTechnicalChallenge/domain/services"

	"testing"

	mocked_repositories "KleverTechnicalChallenge/tests/mocked_repositories"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var upvoteService *services.UpvoteService

func init() {
	godotenv.Overload(".env")
	repository := mocked_repositories.NewUpvoteRepository()
	upvoteService, _ = services.NewUpvoteService(repository)
}

func TestUpvoteService(t *testing.T) {
	upvotes, _ := upvoteService.FindByCommentId("61f46f117942c5844c8cb661")
	numberOfDocumentsBeforeInsertion := len(upvotes)
	newObjID, _ := primitive.ObjectIDFromHex("61f46f117942c5844c8cb661")
	id, _ := upvoteService.Insert(models.Upvote{
		Type:      "upvote",
		CommentId: newObjID,
	})
	upvotes, _ = upvoteService.FindByCommentId("61f46f117942c5844c8cb661")
	numberOfDocumentsAfterInsertion := len(upvotes)
	if numberOfDocumentsBeforeInsertion+1 != numberOfDocumentsAfterInsertion {
		t.Errorf("Erro ao inserir ou ler dados na base de dados")
	}

	upvote, _ := upvoteService.FindById(id)

	if !primitive.IsValidObjectID(id) || len(upvote) == 0 {
		t.Errorf("O id retornado pela insercao e invalido")
	}

	if upvote[0].Type != "upvote" || upvote[0].CommentId != newObjID {
		t.Errorf("Os dados retornados pela insercao sao invalidos")
	}

	upvoteService.DeleteById(id)
	upvote, _ = upvoteService.FindById(id)
	if len(upvote) != 0 {
		t.Errorf("O documento nao foi deletado")
	}
}
