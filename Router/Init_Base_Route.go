package Router

import (
	"Attendance/Controller/Base"
	"github.com/gin-gonic/gin"
)

func Init_Base_Route(rgBase *gin.RouterGroup, BaseContorller *Base.BaseController) {

	rgBase.POST("/login", BaseContorller.Login)
	rgBase.POST("/Application", BaseContorller.Application)
}
