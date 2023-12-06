package router

import (
	"RestApi-Gin/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(Mhscontroller controller.MahasiswaController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hallo Page")
	})

	router.GET("/api/mahasiswa", Mhscontroller.FindAll)
	router.POST("/api/mahasiswa", Mhscontroller.Create)
	router.PATCH("/api/mahasiswa/:mhsId", Mhscontroller.Update)
	router.DELETE("/api/mahasiswa/:mhsId", Mhscontroller.Delete)
	router.GET("/api/mahasiswa/:mhsId", Mhscontroller.FindById)

	return router
}
