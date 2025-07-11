package config

type User struct {
	ID       uint `gorm:"primaryKey"`
	Nama     string
	Email    string
	Password string
	Role     string // admin/user
}
