package service

import (
	"RestApi-Gin/request"
	"RestApi-Gin/response"
)

type MahasiswaService interface {
	Create(mahasiswaRequest request.CreateMahasiswaRequest)
	Update(mahasiswaRequest request.UpdateMahasiswaRequest)
	Delete(mhsId int)
	FindId(mhsId int) response.MahasiswaResponse
	FindAll() []response.MahasiswaResponse
}
