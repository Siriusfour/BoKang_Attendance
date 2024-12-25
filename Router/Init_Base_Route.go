package Router

import (
	"Attendance/Controller/Base"
	"github.com/gin-gonic/gin"
)

func Init_Base_Route(rgBase *gin.RouterGroup) {

	Base_Contorller := Base.NewBase()

	rgBase.POST("/login", Base_Contorller.Login)
	rgBase.POST("/Application", Base_Contorller.Application)
}
