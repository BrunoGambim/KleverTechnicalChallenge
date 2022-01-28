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

func (controller *UpvoteController) GetUpvoteById(ctx context.Context, idDTO *IdDTO) (*GetUpvoteDTO, error) {
	id := idDTO.Id
	upvote, err := controller.upvoteService.FindById(id)
	if err != nil {
		log.Printf(err.Error())
		return &GetUpvoteDTO{}, nil
	}
	if len(upvote) == 0 {
		log.Printf("Not found")
		return &GetUpvoteDTO{}, nil
	}
	response := getAlbumDtoFromModel(upvote[0])
	return response, nil
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

func (controller *UpvoteController) DeleteUpvote(ctx context.Context, idDto *IdDTO) (*empty.Empty, error) {
	err := controller.upvoteService.DeleteById(idDto.Id)
	if err != nil {
		log.Printf(err.Error())
	}
	return &empty.Empty{}, nil
}
