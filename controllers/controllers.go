package controllers

import (
	"MailAoo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func EmailSender(c *gin.Context) {
	EmailFrom := c.PostForm("From")
	Password := c.PostForm("Password")
	EmailTo := c.PostForm("To")
	Subject := c.PostForm("Subject")
	Body := c.PostForm("Body")
	File, err := c.FormFile("File")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cant attach file", "error": err.Error()})
		return
	}
	folderPath := "/home/lenovo/MailApp/attachments"
	UploadFilePath := folderPath + File.Filename
	if err := c.SaveUploadedFile(File, UploadFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't attach the file", "error": err.Error()})
		return
	}
	//Creating dialer for smtp server
	d := gomail.NewDialer(utils.HOST, utils.PORT, EmailFrom, Password)
	m := gomail.NewMessage()
	m.SetHeader("From", EmailFrom)
	m.SetHeader("To", EmailTo)
	m.SetHeader("subject", Subject)
	m.SetBody("text/html", Body)
	m.Attach(UploadFilePath)
	//sending Email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Email send successfully"})
}
