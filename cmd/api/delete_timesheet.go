package api

import (
	"context"
	"errors"

	pb "time_logger/proto-gen/timesheet"
)

func (s *SheetServer) DeleteTimesheet(ctx context.Context, req *pb.DeleteTimesheetRequest) (*pb.DeleteTimesheetResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	// Iterate over the slice to find and delete the timesheet
	for i, ts := range s.Timesheets {
		if ts.Id == req.Id {
			// Remove the timesheet from the slice
			s.Timesheets = append(s.Timesheets[:i], s.Timesheets[i+1:]...)
			return &pb.DeleteTimesheetResponse{Success: true}, nil
		}
	}

	return &pb.DeleteTimesheetResponse{Success: false}, errors.New("timesheet not found with the given ID")
}
