package repository

import (
	"RestApi-Gin/helper"
	"RestApi-Gin/model"
	"RestApi-Gin/request"
	"errors"
	"gorm.io/gorm"
)

type MahasiswaRepositoryImpl struct {
	DB *gorm.DB
}

func NewMahasiswaRepositoryImpl(DB *gorm.DB) *MahasiswaRepositoryImpl {
	return &MahasiswaRepositoryImpl{DB: DB}
}

func (m *MahasiswaRepositoryImpl) Save(mahasiswa model.Mahasiswa) {
	result := m.DB.Create(&mahasiswa)
	helper.PanicError(result.Error)
}

func (m *MahasiswaRepositoryImpl) Update(mahasiswa model.Mahasiswa) {
	var updateMhs = request.UpdateMahasiswaRequest{
		ID:      int(mahasiswa.ID),
		Name:    mahasiswa.Name,
		Address: mahasiswa.Address,
		Jurusan: mahasiswa.Jurusan,
	}
	result := m.DB.Model(&mahasiswa).Updates(updateMhs)
	helper.PanicError(result.Error)
}

func (m *MahasiswaRepositoryImpl) Delete(mhsId int) {
	var mhs model.Mahasiswa
	result := m.DB.Where("id = ?", mhsId).Delete(&mhs)
	helper.PanicError(result.Error)
}

func (m *MahasiswaRepositoryImpl) FindById(mhsId int) (mahasiswa model.Mahasiswa, err error) {
	var mhs model.Mahasiswa
	result := m.DB.Find(&mhs, mhsId)
	if result != nil {
		return mhs, nil
	} else {
		return mhs, errors.New("id not found")
	}
}

func (m *MahasiswaRepositoryImpl) FindAll() []model.Mahasiswa {
	var mhs []model.Mahasiswa
	result := m.DB.Find(&mhs)
	helper.PanicError(result.Error)
	return mhs
}
