package comment_controller

import (
	models "KleverTechnicalChallenge/domain/models"
)

func getCommentDtoFromModel(comment models.Comment) *GetCommentDTO {
	return &GetCommentDTO{
		Id:      comment.Id.Hex(),
		Message: comment.Message,
	}
}

func commentFromCreateDto(commentDto *CreateCommentDTO) models.Comment {
	return models.Comment{
		Message: commentDto.Message,
	}
}
