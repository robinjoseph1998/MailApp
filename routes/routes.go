package routes

import (
	"MailAoo/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRputes(r *gin.Engine) {
	r.POST("/sendemail", controllers.EmailSender)
}
