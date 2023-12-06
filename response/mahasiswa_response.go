package response

type MahasiswaResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Jurusan string `json:"jurusan"`
}
