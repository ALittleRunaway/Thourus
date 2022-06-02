package entrypoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thourus-api/domain/usecase"
)

func GetSpacesInCompany(companyUc *usecase.CompanyUseCase, ctx *gin.Context) {
	companyUid := ctx.Param("uid")

	spaces, err := companyUc.GetSpacesInCompany(companyUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	company, err := companyUc.GetCompanyInfo(companyUid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	data := gin.H{
		"title":       "Spaces in your company",
		"companyName": company.Name,
		"spaces":      spaces,
	}

	ctx.HTML(http.StatusOK, "company.html", data)
}
