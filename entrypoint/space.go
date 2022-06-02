package entrypoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thourus-api/domain/usecase"
)

func GetProjectsInSpace(spaceUc *usecase.SpaceUseCase, ctx *gin.Context) {
	spaceUid := ctx.Param("uid")

	projects, err := spaceUc.GetProjectsInSpace(spaceUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	space, err := spaceUc.GetSpaceInfo(spaceUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	var emptyMessage string
	if len(projects) != 0 {
		emptyMessage = ""
	} else {
		emptyMessage = "There are no projects in the space"
	}

	data := gin.H{
		"title":        "Projects in your company",
		"spaceName":    space.Name,
		"projects":     projects,
		"emptyMessage": emptyMessage,
	}

	ctx.HTML(http.StatusOK, "space.html", data)
}
