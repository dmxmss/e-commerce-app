package entities

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	ParentID *int
	Parent *Category `gorm:"foreignKey:ParentID"`
	Children []Category `gorm:"foreignKey:ParentID"`
}
