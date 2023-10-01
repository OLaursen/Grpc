package main

import (
	"context"
	"log"
	"net"
	

	gRPC "github.com/OLaursen/Grpc/proto"

	"google.golang.org/grpc"
)

type myexercise1Server struct {

	//Obligatory interface for the server type. 
	gRPC.UnimplementedExercise1Server
}



func (s *myexercise1Server) Course(ctx context.Context, courseID *gRPC.CourseID) (*gRPC.CourseInformation, error){

	ack:= &gRPC.CourseInformation{
		CourseID:   courseID,
		CourseName: "Distributed Systems",
	}
	return ack , nil; 

}

func main() {

	//Listener on localhost:5400
	list, err:= net.Listen("tcp", "localhost:5400");
	if err != nil {
		log.Fatalf("Failed to listen on port 5400: %v", err)
	}

	//Create server
	grpcServer := grpc.NewServer();
	//Create instance of server struct
	server := &myexercise1Server{
		//Fields go here
	};
	//Register server
	gRPC.RegisterExercise1Server(grpcServer, server);

	grpcServer.Serve(list)
	//Anything below won't compute, as Serve is a blocking call. 
	
}