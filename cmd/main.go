package main

import (
	"log"
	"net"
	"time_logger/cmd/api"
	pb "time_logger/proto-gen/timesheet"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTimesheetServiceServer(grpcServer, &api.SheetServer{
		Timesheets: make(map[string]*pb.Timesheet),
		IdCounter:  0,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
