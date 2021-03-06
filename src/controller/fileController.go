package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

//UploadFile TODO: Fix Lint Error
type UploadFile struct {
}

//UploadSingleFile TODO: Fix Lint Error
func (fileUploader UploadFile) UploadSingleFile(ctx *gin.Context) {
	request := ctx.Request
	file, handler, err := request.FormFile("ex1")
	if err != nil {
		log.Error("Error Retrieving the File")
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	f, err := os.OpenFile("uploadData/"+handler.Filename, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Error Opening the File", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer f.Close()
	io.Copy(f, file)

	log.WithFields(log.Fields{
		"FileName": handler.Filename,
		"Filesize": handler.Size,
	}).Info("File Uploaded Successfully")
	ctx.JSON(200, gin.H{"home": "sweet home"})
}

//Health TODO: Fix Lint Error
func (fileUploader UploadFile) Health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Status": "Healthy"})
}

//NewFileUploader TODO: Fix Lint Error
func NewFileUploader() UploadFile {
	return UploadFile{}
}
