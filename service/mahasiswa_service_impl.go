package service

import (
	"RestApi-Gin/helper"
	"RestApi-Gin/model"
	"RestApi-Gin/repository"
	"RestApi-Gin/request"
	"RestApi-Gin/response"
	"github.com/go-playground/validator/v10"
)

type MahasiswaServiceImpl struct {
	MahasiswaRepository repository.MahasiswaRepository
	Validate            *validator.Validate
}

func NewMahasiswaServiceImpl(mahasiswaRepository repository.MahasiswaRepository, validate *validator.Validate) *MahasiswaServiceImpl {
	return &MahasiswaServiceImpl{
		MahasiswaRepository: mahasiswaRepository,
		Validate:            validate,
	}
}

func (m *MahasiswaServiceImpl) Create(mahasiswaRequest request.CreateMahasiswaRequest) {
	err := m.Validate.Struct(mahasiswaRequest)
	helper.PanicError(err)
	mahasiswaModel := model.Mahasiswa{
		Name:    mahasiswaRequest.Name,
		Address: mahasiswaRequest.Address,
		Jurusan: mahasiswaRequest.Jurusan,
	}
	m.MahasiswaRepository.Save(mahasiswaModel)
}

func (m *MahasiswaServiceImpl) Update(mahasiswaRequest request.UpdateMahasiswaRequest) {
	err2 := m.Validate.Struct(mahasiswaRequest)
	helper.PanicError(err2)
	mahasiswa, err := m.MahasiswaRepository.FindById(mahasiswaRequest.ID)
	helper.PanicError(err)

	mahasiswa.Name = mahasiswaRequest.Name
	mahasiswa.Address = mahasiswaRequest.Address
	mahasiswa.Jurusan = mahasiswaRequest.Jurusan
	m.MahasiswaRepository.Update(mahasiswa)
}

func (m *MahasiswaServiceImpl) Delete(mhsId int) {
	m.MahasiswaRepository.Delete(mhsId)
}

func (m *MahasiswaServiceImpl) FindId(mhsId int) response.MahasiswaResponse {
	mahasiswa, err := m.MahasiswaRepository.FindById(mhsId)
	helper.PanicError(err)

	mahasiswaResponse := response.MahasiswaResponse{
		ID:      int(mahasiswa.ID),
		Name:    mahasiswa.Name,
		Address: mahasiswa.Address,
		Jurusan: mahasiswa.Jurusan,
	}

	return mahasiswaResponse
}

func (m *MahasiswaServiceImpl) FindAll() []response.MahasiswaResponse {
	result := m.MahasiswaRepository.FindAll()

	var responses []response.MahasiswaResponse
	for _, value := range result {
		respon := response.MahasiswaResponse{
			ID:      int(value.ID),
			Name:    value.Name,
			Address: value.Address,
			Jurusan: value.Jurusan,
		}
		responses = append(responses, respon)
	}
	return responses
}
