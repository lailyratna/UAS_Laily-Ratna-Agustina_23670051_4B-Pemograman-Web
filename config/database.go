package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ===== Struct Model =====
type User struct {
	ID       uint `gorm:"primaryKey"`
	Nama     string
	Email    string
	Password string
	Role     string
}

type Lab struct {
	ID        uint `gorm:"primaryKey"`
	NamaLab   string
	Lokasi    string
	Kapasitas int
}

type Alat struct {
	ID       uint `gorm:"primaryKey"`
	NamaAlat string
	Jumlah   int
	Kondisi  string
	LabID    uint
	Lab      Lab // relasi untuk preload
}

// ===== Koneksi DB dan Migrasi =====
func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/peminjaman_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal koneksi database: " + err.Error())
	}
	fmt.Println("Database terkoneksi")

	// AutoMigrate model
	DB.AutoMigrate(&User{}, &Lab{}, &Alat{})
}
