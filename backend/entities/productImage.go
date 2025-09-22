package entities

type ProductImage struct {
	ID int `gorm:"primaryKey"`
	ProductID int `gorm:"index;not null"`
	URL string `gorm:"type:text;not null"`
	Position uint `gorm:"not null"`
}
