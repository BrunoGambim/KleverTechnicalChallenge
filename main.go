package main

import (
	"fmt"
	"log"
	"net"
	"os"

	comment_controller "KleverTechnicalChallenge/controllers/comment_controller"
	upvote_controller "KleverTechnicalChallenge/controllers/upvote_controller"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func getNetListener() net.Listener {
	port := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(fmt.Sprintf(err.Error()))
	}

	return lis
}

func loadEnvFiles() {
	err := godotenv.Overload(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

}

func main() {
	loadEnvFiles()

	list := getNetListener()
	grpcServer := grpc.NewServer()

	comment_controller.RegisterCommentControllerServer(grpcServer, comment_controller.NewCommentController())
	upvote_controller.RegisterUpvoteControllerServer(grpcServer, upvote_controller.NewUpvoteController())

	err := grpcServer.Serve(list)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
