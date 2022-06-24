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

func DeleteProject(projectUc *usecase.ProjectUseCase, ctx *gin.Context) {
	projectUid := ctx.Param("uid")

	err := projectUc.DeleteProject(projectUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your project has been successfully deleted.",
	})
}

func AddProject(projectUc *usecase.ProjectUseCase, spaceUc *usecase.SpaceUseCase, ctx *gin.Context) {
	projectName := ctx.Query("project_name")

	spaceUid, err := ctx.Request.Cookie("space_uid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get space uid",
		})
		return
	}

	space, err := spaceUc.GetSpaceInfo(spaceUid.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	project, err := projectUc.AddProject(projectName, space.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your space has been successfully added.",
		"project": project,
	})
}
