package request

type UpdateMahasiswaRequest struct {
	ID      int    `validate:"required" json:"id"`
	Name    string `validate:"required,min=1,max=255" json:"name"`
	Address string `validate:"required,min=1,max=255" json:"address"`
	Jurusan string `validate:"required,min=1,max=255" json:"jurusan"`
}
