package request

type CreateMahasiswaRequest struct {
	Name    string `validate:"required,min=1,max=255" json:"name"`
	Address string `validate:"required,min=1,max=255" json:"address"`
	Jurusan string `validate:"required,min=1,max=255" json:"jurusan"`
}
