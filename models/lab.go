package config

type Lab struct {
	ID        uint `gorm:"primaryKey"`
	NamaLab   string
	Lokasi    string
	Kapasitas int
}
