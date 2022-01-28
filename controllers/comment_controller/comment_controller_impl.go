package comment_controller

import (
	"context"
	"log"

	services "KleverTechnicalChallenge/domain/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type CommentController struct {
	UnimplementedCommentControllerServer
	commentService *services.CommentService
}

func NewCommentController() *CommentController {
	service, err := services.NewCommentService()
	if err != nil {
		log.Fatal(err)
	}

	return &CommentController{
		commentService: service,
	}
}

func (controller *CommentController) GetAllComments(e *empty.Empty, stream CommentController_GetAllCommentsServer) error {
	commentList, err := controller.commentService.FindAll()
	if err != nil {
		log.Printf(err.Error())
		return nil
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
		log.Printf(err.Error())
		return &empty.Empty{}, nil
	}
	return &empty.Empty{}, nil
}
