package Router

import (
	"Attendance/Controller/Base"
	"Attendance/Controller/Leader"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	//路由分组

	rgAuth := r.Group("Attendance/Api/Leader")
	rgBase := r.Group("Attendance/Api/Base")

	//注册所有组别的路由
	init_Base_Paltform_Router(rgAuth, rgBase)
}

func init_Base_Paltform_Router(rgAuth *gin.RouterGroup, rgBase *gin.RouterGroup) {

	BaseController := Base.NewBase()
	LeaderController := Leader.NewLeader(BaseController)

	Init_Leader_Route(rgAuth, LeaderController)
	Init_Base_Route(rgBase, BaseController)

}
