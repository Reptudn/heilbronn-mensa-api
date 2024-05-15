package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenu(ctx *gin.Context) {
	lang := ctx.Query("lang")
	date := ctx.Query("date")
	menu := GetMenuData(date, lang)
	ctx.IndentedJSON(http.StatusOK, menu)
}

func TestHi(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "hello")
}

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/mensa/menu/:lang/:date", GetMenu)
	r.GET("/", TestHi)
}
