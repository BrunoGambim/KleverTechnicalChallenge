package main

import (
	"fmt"
	"log"
	"net"

	comment_controller "KleverTechnicalChallenge/controllers/comment_controller"
	upvote_controller "KleverTechnicalChallenge/controllers/upvote_controller"

	"google.golang.org/grpc"
)

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	return lis
}

func main() {
	list := getNetListener(9000)
	grpcServer := grpc.NewServer()

	comment_controller.RegisterCommentControllerServer(grpcServer, comment_controller.NewCommentController())
	upvote_controller.RegisterUpvoteControllerServer(grpcServer, upvote_controller.NewUpvoteController())

	err := grpcServer.Serve(list)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
