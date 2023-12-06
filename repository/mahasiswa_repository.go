package repository

import "RestApi-Gin/model"

type MahasiswaRepository interface {
	Save(mahasiswa model.Mahasiswa)
	Update(mahasiswa model.Mahasiswa)
	Delete(mhsId int)
	FindById(mhsId int) (mahasiswa model.Mahasiswa, err error)
	FindAll() []model.Mahasiswa
}
