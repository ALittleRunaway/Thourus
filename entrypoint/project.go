package entrypoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thourus-api/domain/usecase"
)

func GetDocumentsInProject(projectUc *usecase.ProjectUseCase, ctx *gin.Context) {
	projectUid := ctx.Param("uid")

	documents, err := projectUc.GetDocumentsInProject(projectUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	project, err := projectUc.GetProjectInfo(projectUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	var emptyMessage string
	if len(documents) != 0 {
		emptyMessage = ""
	} else {
		emptyMessage = "There are no documents in the project"
	}

	data := gin.H{
		"title":        "Documents in your project",
		"projectName":  project.Name,
		"documents":    documents,
		"emptyMessage": emptyMessage,
	}

	ctx.HTML(http.StatusOK, "project.html", data)
}
