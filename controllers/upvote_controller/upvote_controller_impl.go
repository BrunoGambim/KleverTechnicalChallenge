package upvote_controller

import (
	"context"
	"log"

	services "KleverTechnicalChallenge/domain/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type UpvoteController struct {
	UnimplementedUpvoteControllerServer
	upvoteService *services.UpvoteService
}

func NewUpvoteController() *UpvoteController {
	service, err := services.NewUpvoteService()
	if err != nil {
		log.Fatal(err)
	}

	return &UpvoteController{
		upvoteService: service,
	}
}

func (controller *UpvoteController) CreateUpvote(ctx context.Context, upvoteDto *CreateUpvoteDTO) (*empty.Empty, error) {
	upvote, err := upvoteFromCreateDto(upvoteDto)
	if err != nil {
		log.Printf(err.Error())
		return &empty.Empty{}, nil
	}

	_, err = controller.upvoteService.Insert(upvote)
	if err != nil {
		log.Printf(err.Error())
		return &empty.Empty{}, nil
	}
	return &empty.Empty{}, nil
}
