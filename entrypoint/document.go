package entrypoint

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"thourus-api/domain/usecase"
)

func UploadNewDocument(documentUc *usecase.DocumentUseCase, ctx *gin.Context) {

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	ctx.JSON(http.StatusOK, gin.H{"filepath": filepath})

	//companyUid := ctx.Param("uid")
	//spaces, err := companyUc.GetSpacesInCompany(companyUid)
	//if err != nil {
	//	fmt.Println(err)
	//	ctx.JSON(http.StatusInternalServerError, gin.Error{Err: err}.Error())
	//}
	//
	//data := gin.H{
	//	"title":  "Spaces in your company",
	//	"spaces": spaces,
	//}
	//
	//ctx.HTML(http.StatusOK, "company.html", data)
}
