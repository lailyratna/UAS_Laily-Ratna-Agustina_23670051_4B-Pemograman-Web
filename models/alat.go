package config

type Alat struct {
	ID       uint `gorm:"primaryKey"`
	NamaAlat string
	Jumlah   int
	Kondisi  string
	LabID    uint
}
