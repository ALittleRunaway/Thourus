package entrypoint

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strings"
	"thourus-api/domain/usecase"
)

type BindFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func UploadNewDocument(documentUc *usecase.DocumentUseCase, mailUc *usecase.MailUseCase, cacheUc *usecase.CacheUseCase, ctx *gin.Context) {
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

	document, err := documentUc.UploadNewDocument(file, userUid.Value, projectUid.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = cacheUc.SaveDocumentVersion(document.Uid, "1")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//err = mailUc.SendUpdates(file.Filename)
	//if err != nil {
	//	return
	//}
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

func UpdateDocument(documentUc *usecase.DocumentUseCase, mailUc *usecase.MailUseCase, cacheUc *usecase.CacheUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")

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

	version, err := cacheUc.GetDocumentVersion(documentUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = documentUc.UpdateDocument(file, userUid.Value, documentUid, version)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = cacheUc.IncrementDocumentVersion(documentUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = mailUc.SendUpdates(file.Filename)
	if err != nil {
		return
	}
	// TODO: add hierarchy
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})

}

func AddDocument(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add_document.html", struct{}{})
}

func UpdateDocumentView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "update_document.html", struct{}{})
}

func ShowHistory(documentUc *usecase.DocumentUseCase, cacheUC *usecase.CacheUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")

	documentHistoryRows, err := documentUc.GetDocumentHistory(documentUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}
	documentVersion, err := cacheUC.GetDocumentVersion(documentUid)
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
		"history":         documentHistoryRows,
		"emptyMessage":    emptyMessage,
		"document":        documentHistoryRows[0].Document,
		"documentVersion": documentVersion,
	}

	ctx.HTML(http.StatusOK, "history.html", data)
}

func DownloadDocument(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {
	documentUid := ctx.Param("uid")
	version := ctx.Query("version")

	document, err := documentUc.DownloadDocument(documentUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var path string
	if version != "" {
		nameExtention := strings.Split(document.Name, ".")
		path = fmt.Sprintf("%s-v%v.%s", nameExtention[0], version, nameExtention[1])
	} else {
		path = document.Name
	}

	ctx.FileAttachment(document.Path, path)
}
