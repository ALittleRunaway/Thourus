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

func UploadNewDocument(documentUc *usecase.DocumentUseCase, mailUc *usecase.MailUseCase, ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}
	userUid, err := ctx.Request.Cookie("user_uid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get user uid",
		})
		return
	}
	//projectUid, err := ctx.Request.Cookie("project_uid")
	//if err != nil {
	//	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	//		"message": "Unable to get project uid",
	//	})
	//	return
	//}
	projectUid := "b8c2c6850b32"

	err = documentUc.UploadNewDocument(file, userUid.Value, projectUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = mailUc.SendUpdates()
	if err != nil {
		return
	}
	// TODO: add hierarchy
	// File saved successfully. Return proper result
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

}

func DeleteDocument(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")

	err := documentUc.DeleteDocument(documentUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully deleted.",
	})

}

func AddDocument(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_document.html", struct{}{})
}
