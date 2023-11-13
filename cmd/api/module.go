package api

import (
	"sync"
	pb "time_logger/proto-gen/timesheet"
)

type SheetServer struct {
	pb.UnimplementedTimesheetServiceServer
	Timesheets map[string]*pb.Timesheet
	Mu         sync.Mutex // Protects timesheets
	IdCounter  int64
}
