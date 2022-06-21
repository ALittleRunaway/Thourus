package entrypoint

import "C"
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

	projectUid, err := ctx.Request.Cookie("project_uid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get project uid",
		})
		return
	}

	err = documentUc.UploadNewDocument(file, userUid.Value, projectUid.Value)
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

func ShowHistory(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")

	documentHistoryRows, err := documentUc.GetDocumentHistory(documentUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	var emptyMessage string
	if len(documentHistoryRows) != 0 {
		emptyMessage = ""
	} else {
		emptyMessage = "There are no history about this document"
	}

	data := gin.H{
		"history":      documentHistoryRows,
		"emptyMessage": emptyMessage,
		"document":     documentHistoryRows[0].Document,
	}

	ctx.HTML(http.StatusOK, "history.html", data)
}

func DownloadDocument(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")

	document, err := documentUc.DownloadDocument(documentUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.FileAttachment(document.Path, document.Name)
}
