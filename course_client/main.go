package main

import (
	"context"
	"fmt"
	"log"
	"time"

	gRPC "github.com/OLaursen/Grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	// Establish a connection to the gRPC server.
	conn, err := grpc.Dial("localhost:5400", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client instance.
	client := gRPC.NewExercise1Client(conn)

	// Create a context with a timeout for the RPC call.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the "Course" RPC method.
	courseID := &gRPC.CourseID{CourseID: 1}
	response, err := client.Course(ctx, courseID)
	if err != nil {
		log.Fatalf("Error calling Course: %v", err)
	}

	// Print the response from the server.
	fmt.Printf("Course Name: %s\n", response.CourseName)
	fmt.Printf("Course ID: %d\n", response.CourseID.CourseID)
}