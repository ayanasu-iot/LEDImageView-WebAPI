package controllers

import (
	"LedImageView-WebAPI/pkg/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func errorResponse(err error, ctx *gin.Context, status int) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.JSON(status, err.Error())
	log.Fatalln(err.Error())
}

func UploadAnimation(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		errorResponse(err, ctx, http.StatusBadRequest)
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	fileName := filepath.Base(header.Filename)
	_, _, err = image.DecodeConfig(buf)
	if err != nil {
		errorResponse(err, ctx, http.StatusBadRequest)
	}
	dstPath := "./uploads/" + fileName
	err = ctx.SaveUploadedFile(header, dstPath)
	if err != nil {
		errorResponse(err, ctx, http.StatusInternalServerError)
	}
	data := models.StatusData{
		FileName: fileName,
		Status:   http.StatusOK,
	}
	status, err := json.Marshal(data)
	if err != nil {
		errorResponse(err, ctx, http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, status)
	ShowImage(dstPath)
}