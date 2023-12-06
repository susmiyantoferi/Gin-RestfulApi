package controller

import (
	"RestApi-Gin/helper"
	"RestApi-Gin/request"
	"RestApi-Gin/response"
	"RestApi-Gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MahasiswaControllerImpl struct {
	MahasiswaService service.MahasiswaService
}

func NewMahasiswaControllerImpl(mahasiswaService service.MahasiswaService) *MahasiswaControllerImpl {
	return &MahasiswaControllerImpl{MahasiswaService: mahasiswaService}
}

func (m *MahasiswaControllerImpl) Create(ctx *gin.Context) {
	createRequest := request.CreateMahasiswaRequest{}
	err := ctx.ShouldBindJSON(&createRequest)
	helper.PanicError(err)

	m.MahasiswaService.Create(createRequest)
	responseWeb := response.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   createRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSONP(http.StatusCreated, responseWeb)
}

func (m *MahasiswaControllerImpl) Update(ctx *gin.Context) {
	updateRequest := request.UpdateMahasiswaRequest{}
	err := ctx.ShouldBindJSON(&updateRequest)
	helper.PanicError(err)

	mhsId := ctx.Param("mhsId")
	id, err := strconv.Atoi(mhsId)
	helper.PanicError(err)
	updateRequest.ID = id
	m.MahasiswaService.Update(updateRequest)

	responseWeb := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   updateRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSONP(http.StatusOK, responseWeb)

}

func (m *MahasiswaControllerImpl) Delete(ctx *gin.Context) {
	mhsId := ctx.Param("mhsId")
	id, err := strconv.Atoi(mhsId)
	helper.PanicError(err)
	m.MahasiswaService.Delete(id)

	responseWeb := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSONP(http.StatusOK, responseWeb)
}

func (m *MahasiswaControllerImpl) FindAll(ctx *gin.Context) {
	allData := m.MahasiswaService.FindAll()

	responseWeb := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   allData,
	}

	ctx.JSON(http.StatusOK, responseWeb)
}

func (m *MahasiswaControllerImpl) FindById(ctx *gin.Context) {
	mhsId := ctx.Param("mhsId")
	id, err := strconv.Atoi(mhsId)
	helper.PanicError(err)
	data := m.MahasiswaService.FindId(id)

	responseWeb := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, responseWeb)
}
