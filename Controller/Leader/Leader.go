package Leader

import (
	"Attendance/Controller/Base"
	"Attendance/Server/LeaderServer"
)

type LeaderController struct {
	BaseController *Base.Base
	LeaderServer   *LeaderServer.LeaderServer
}

func NewLeader(base *Base.Base) *LeaderController {
	return &LeaderController{
		BaseController: base,
		LeaderServer:   LeaderServer.New_Leader_Server(),
	}

}
