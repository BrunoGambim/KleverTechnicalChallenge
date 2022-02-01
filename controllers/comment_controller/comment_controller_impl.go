package comment_controller

import (
	"context"
	"log"

	controllers_utils "KleverTechnicalChallenge/controllers/utils"
	comment_repository "KleverTechnicalChallenge/database/repositories/comment_repository"
	services "KleverTechnicalChallenge/domain/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type CommentController struct {
	UnimplementedCommentControllerServer
	commentService *services.CommentService
}

func NewCommentController() *CommentController {
	repository, err := comment_repository.NewCommentRepository()
	if err != nil {
		log.Fatalf(err.Error())
	}

	service, err := services.NewCommentService(repository)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &CommentController{
		commentService: service,
	}
}

func (controller *CommentController) GetAllComments(e *empty.Empty, stream CommentController_GetAllCommentsServer) error {
	commentList, err := controller.commentService.FindAll()
	if err != nil {
		return controllers_utils.Handle(err)
	}

	for _, comment := range commentList {
		stream.Send(getCommentDtoFromModel(comment))
	}
	return nil
}

func (controller *CommentController) CreateComment(ctx context.Context, commentDto *CreateCommentDTO) (*empty.Empty, error) {
	comment := commentFromCreateDto(commentDto)
	_, err := controller.commentService.Insert(comment)
	if err != nil {
		return &empty.Empty{}, controllers_utils.Handle(err)
	}

	return &empty.Empty{}, nil
}
