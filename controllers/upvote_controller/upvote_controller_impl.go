package upvote_controller

import (
	"context"
	"log"

	controllers_utils "KleverTechnicalChallenge/controllers/utils"
	upvote_repository "KleverTechnicalChallenge/database/repositories/upvote_repository"
	services "KleverTechnicalChallenge/domain/services"

	"github.com/golang/protobuf/ptypes/empty"
)

type UpvoteController struct {
	UnimplementedUpvoteControllerServer
	upvoteService *services.UpvoteService
}

func NewUpvoteController() *UpvoteController {
	repository, err := upvote_repository.NewUpvoteRepository()
	if err != nil {
		log.Fatal(err)
	}

	service, err := services.NewUpvoteService(repository)
	if err != nil {
		log.Fatal(err)
	}

	return &UpvoteController{
		upvoteService: service,
	}
}

func (controller *UpvoteController) GetUpvoteById(ctx context.Context, idDto *IdDTO) (*GetUpvoteDTO, error) {
	id, err := idFromDto(idDto)
	if err != nil {
		return &GetUpvoteDTO{}, controllers_utils.HandleInvalidIdError(err)
	}
	upvote, err := controller.upvoteService.FindById(id)
	if err != nil {
		return &GetUpvoteDTO{}, controllers_utils.HandleUnknownError(err)
	}
	if len(upvote) == 0 {
		return &GetUpvoteDTO{}, controllers_utils.HandleNotFoundError()
	}
	response := getAlbumDtoFromModel(upvote[0])
	return response, nil
}

func (controller *UpvoteController) GetUpvotesByCommentId(idDto *IdDTO, stream UpvoteController_GetUpvotesByCommentIdServer) error {
	id, err := idFromDto(idDto)
	if err != nil {
		return controllers_utils.HandleInvalidIdError(err)
	}
	upvotes, err := controller.upvoteService.FindByCommentId(id)
	if err != nil {
		return controllers_utils.HandleUnknownError(err)
	}
	for _, upvote := range upvotes {
		stream.Send(getAlbumDtoFromModel(upvote))
	}
	return nil
}

func (controller *UpvoteController) CreateUpvote(ctx context.Context, upvoteDto *CreateUpvoteDTO) (*empty.Empty, error) {
	upvote, err := upvoteFromCreateDto(upvoteDto)
	if err != nil {
		return &empty.Empty{}, controllers_utils.HandleInvalidIdError(err)
	}

	_, err = controller.upvoteService.Insert(upvote)
	if err != nil {
		return &empty.Empty{}, controllers_utils.HandleUnknownError(err)
	}
	return &empty.Empty{}, nil
}

func (controller *UpvoteController) DeleteUpvote(ctx context.Context, idDto *IdDTO) (*empty.Empty, error) {
	id, err := idFromDto(idDto)
	if err != nil {
		return &empty.Empty{}, controllers_utils.HandleInvalidIdError(err)
	}
	err = controller.upvoteService.DeleteById(id)
	if err != nil {
		return &empty.Empty{}, controllers_utils.HandleUnknownError(err)
	}
	return &empty.Empty{}, nil
}
