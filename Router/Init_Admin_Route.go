package Router

import "github.com/gin-gonic/gin"
import "Attendance/Controller/Leader"

func Init_Leader_Route(rgLeader *gin.RouterGroup, LeaderController *Leader.LeaderController) {

	rgLeader.POST("/view", LeaderController.View)
}
