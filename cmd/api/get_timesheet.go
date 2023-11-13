package api

import (
	"context"
	pb "time_logger/proto-gen/timesheet"
)

func (s *SheetServer) GetTimesheetsByContractor(ctx context.Context, req *pb.GetTimesheetsByContractorRequest) (*pb.GetTimesheetsResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	var contractorTimesheets []*pb.Timesheet
	for _, ts := range s.Timesheets {
		if ts.ContractorId == req.ContractorId {
			contractorTimesheets = append(contractorTimesheets, ts)
		}
	}

	return &pb.GetTimesheetsResponse{
		Timesheets: contractorTimesheets,
	}, nil
}

func (s *SheetServer) GetTimesheetsByContractorAndClient(ctx context.Context, req *pb.GetTimesheetsByContractorAndClientRequest) (*pb.GetTimesheetsResponse, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	var contractorClientTimesheets []*pb.Timesheet
	for _, ts := range s.Timesheets {
		if ts.ContractorId == req.ContractorId && ts.ClientId == req.ClientId {
			contractorClientTimesheets = append(contractorClientTimesheets, ts)
		}
	}

	return &pb.GetTimesheetsResponse{
		Timesheets: contractorClientTimesheets,
	}, nil
}
