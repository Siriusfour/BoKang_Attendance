package Router

import (
	"Attendance/Controller/Base"
	"github.com/gin-gonic/gin"
)

func Init_Base_Route(rgBase *gin.RouterGroup) {

	Base_Contorller := Base.NewBase()

	rgBase.PUT("/login", Base_Contorller.Login)
}
