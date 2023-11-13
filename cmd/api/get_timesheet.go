package api

import (
	"context"
	pb "time_logger/proto-gen/timesheet"
)

func (s *SheetServer) GetTimesheetsByContractor(ctx context.Context, req *pb.GetTimesheetsByContractorRequest) (*pb.GetTimesheetsByContractorResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	var contractorTimesheets []*pb.Timesheet
	for _, ts := range s.Timesheets {
		if ts.ContractorId == req.ContractorId {
			contractorTimesheets = append(contractorTimesheets, ts)
		}
	}

	return &pb.GetTimesheetsByContractorResponse{
		Timesheets: contractorTimesheets,
	}, nil
}
