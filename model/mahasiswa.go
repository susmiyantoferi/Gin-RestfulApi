package model

type Mahasiswa struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"notnull" json:"name"`
	Address string `gorm:"notnull" json:"address"`
	Jurusan string `gorm:"notnull" json:"jurusan"`
}
