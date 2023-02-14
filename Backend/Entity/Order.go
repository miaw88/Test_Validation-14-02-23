package entity

import(

	"gorm.io/gorm"
	// "github.com/asaskevich/govalidator"
)

type Video struct {
	gorm.Model
	Name string `valid:"required"`
	Url string `gorm:"uniqueIndex" valid:"url"`
}