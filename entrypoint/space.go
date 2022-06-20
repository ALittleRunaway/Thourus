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

func DeleteSpace(spaceUc *usecase.SpaceUseCase, ctx *gin.Context) {
	spaceUid := ctx.Param("uid")

	err := spaceUc.DeleteSpace(spaceUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your space has been successfully deleted.",
	})
}

func AddSpace(spaceUc *usecase.SpaceUseCase, companyUc *usecase.CompanyUseCase, ctx *gin.Context) {
	spaceName := ctx.Query("space_name")

	companyUid, err := ctx.Request.Cookie("company_uid")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get company uid",
		})
		return
	}

	company, err := companyUc.GetCompanyInfo(companyUid.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	space, err := spaceUc.AddSpace(company.Id, spaceName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your space has been successfully added.",
		"space":   space,
	})
}
