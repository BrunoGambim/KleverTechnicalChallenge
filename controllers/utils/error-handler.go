package controllers_utils

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Handle(err error) error {
	if errors.Is(err, mongo.ErrNoDocuments) {
		return handleNotFoundError(err)
	} else if errors.Is(err, primitive.ErrInvalidHex) {
		return handleInvalidIdError(err)
	}
	return handleUnknownError(err)
}

func handleUnknownError(err error) error {
	log.Printf(err.Error())
	return status.Error(codes.Unknown, "Unknown error")
}

func handleNotFoundError(err error) error {
	log.Printf(err.Error())
	return status.Error(codes.NotFound, "Not found")
}

func handleInvalidIdError(err error) error {
	log.Printf(err.Error())
	return status.Error(codes.InvalidArgument, "Invalid id")
}
