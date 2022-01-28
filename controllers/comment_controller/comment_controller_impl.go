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

func (controller *CommentController) GetAllComments(ctx context.Context, e *empty.Empty) (*GetAllCommentDTO, error) {
	commentList, err := controller.commentService.FindAll()
	if err != nil {
		log.Printf(err.Error())
		return &GetAllCommentDTO{}, nil
	}
	response := GetAllCommentDTO{}
	for _, comment := range commentList {
		response.Comments = append(response.Comments, getCommentDtoFromModel(comment))
	}
	return &response, nil
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
