package Leader

import (
	"Attendance/Controller/Base"
	"Attendance/Server/LeaderServer"
)

type LeaderController struct {
	BaseController *Base.BaseController
	LeaderServer   *LeaderServer.LeaderServer
}

func NewLeader(base *Base.BaseController) *LeaderController {
	return &LeaderController{
		BaseController: base,
		LeaderServer:   LeaderServer.New_Leader_Server(),
	}

}
