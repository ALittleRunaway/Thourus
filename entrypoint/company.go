package entrypoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"thourus-api/domain/usecase"
)

func GetSpacesInCompany(companyUc *usecase.CompanyUseCase, ctx *gin.Context) {
	companyUid := ctx.Param("uid")
	spaces, err := companyUc.GetSpacesInCompany(companyUid)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	}

	data := gin.H{
		"title":  "Spaces in your company",
		"spaces": spaces,
	}

	ctx.HTML(http.StatusOK, "index.html", data)
}
