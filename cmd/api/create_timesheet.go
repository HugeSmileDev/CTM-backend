package api

import (
	"context"
	"fmt"
	"time"

	pb "time_logger/proto-gen/timesheet"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *SheetServer) CreateTimesheet(ctx context.Context, req *pb.CreateTimesheetRequest) (*pb.Timesheet, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	fmt.Printf("Create Timesheet request came from \"%s\"\n", req.ClientId)

	// Check if a timesheet with the same details already exists
	for _, ts := range s.Timesheets {
		if ts.ContractorId == req.ContractorId && ts.ClientId == req.ClientId && ts.Date == req.Date {
			return nil, fmt.Errorf("a timesheet for this contractor, client, and date already exists")
		}
	}

	// Parse the date to get the day of the week
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid date format: %v", err)
	}
	dayOfWeek := parsedDate.Weekday().String()

	var hourlyRate float32 = 35
	var nonBillableHours float32 = 0
	totalWorkHours := req.TotalHours - nonBillableHours
	pay := hourlyRate * totalWorkHours

	// Create a new Timesheet from the request
	timesheet := &pb.Timesheet{
		Id:               s.generateID(),
		ContractorId:     req.ContractorId,
		ClientId:         req.ClientId,
		Date:             req.Date,
		DayOfWeek:        dayOfWeek,
		Task:             req.Task,
		WorkDetails:      req.WorkDetails,
		TotalHours:       req.TotalHours,
		NonBillableHours: nonBillableHours,
		TotalWorkHours:   totalWorkHours,
		HourlyRate:       hourlyRate,
		Pay:              pay,
	}

	// Store the timesheet in the in-memory map
	s.Timesheets[timesheet.Id] = timesheet

	return timesheet, nil
}

func (s *SheetServer) generateID() string {
	s.IdCounter++
	return fmt.Sprintf("ts_%d", s.IdCounter)
}
