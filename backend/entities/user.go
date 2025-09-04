package entities

type User struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Admin bool `gorm:"default:false"`
}
