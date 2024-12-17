package Router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {

	//路由分组
	rgPublic := r.Group("Attendance/Api/staff")
	rgAuth := r.Group("Attendance/Api/admin")
	rgBase := r.Group("Attendance/Api/Base")

	//注册所有组别的路由
	init_Base_Paltform_Router(rgPublic, rgAuth, rgBase)
}

func init_Base_Paltform_Router(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup, rgBase *gin.RouterGroup) {

	Init_Admin_Route(rgAuth)
	Init_worker_Route(rgPublic)
	Init_Base_Route(rgBase)

}
