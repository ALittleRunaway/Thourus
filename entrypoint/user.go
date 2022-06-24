package entrypoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thourus-api/domain/usecase"
)

func LoginUser(userUc *usecase.UserUseCase, ctx *gin.Context) {
	userLogin := ctx.Query("login")
	userPassword := ctx.Query("pass")

	user, err := userUc.LoginUser(userPassword, userLogin)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", struct{}{})
}
