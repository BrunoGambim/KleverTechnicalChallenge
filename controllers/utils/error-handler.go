package controllers_utils

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleUnknownError(err error) error {
	log.Printf(err.Error())
	return status.Error(codes.Unknown, "Unknown error")
}

func HandleNotFoundError() error {
	log.Printf("Not found")
	return status.Error(codes.NotFound, "Not found")
}

func HandleInvalidIdError(err error) error {
	log.Printf(err.Error())
	return status.Error(codes.InvalidArgument, "Invalid id")
}
