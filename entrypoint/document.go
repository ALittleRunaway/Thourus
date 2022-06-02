package entrypoint

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"thourus-api/domain/usecase"
)

type BindFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func UploadNewDocument(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	err = documentUc.UploadNewDocument(file)
	if err != nil {
		return
	}

	// File saved successfully. Return proper result
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

func AddDocument(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_document.html", struct{}{})
}
