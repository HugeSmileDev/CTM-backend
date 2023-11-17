package api

import (
	"context"
	"errors"
	"fmt"

	pb "time_logger/proto-gen/timesheet"
)

func (s *SheetServer) UpdateTimesheet(ctx context.Context, req *pb.UpdateTimesheetRequest) (*pb.Timesheet, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	updatedTimesheet := req.Timesheet

	fmt.Printf("Timesheet ID: %s\n", req.Timesheet.GetId())

	// Iterate over the slice to find and update the timesheet
	for i, ts := range s.Timesheets {
		if ts.Id == updatedTimesheet.Id {
			s.Timesheets[i] = updatedTimesheet
			return updatedTimesheet, nil
		}
	}

	return nil, errors.New("timesheet not found with the given ID")
}
