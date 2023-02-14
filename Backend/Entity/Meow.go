package entity

import(
	"gorm.io/gorm"

)

type Meow struct{
	gorm.Model
	Test_Matches string `valid:"maxstringlength(10),required~Test cannot be blank"`
	Phone string `valid:"length(0|10),matches([0-9])"`
	Cost int `valid:"range(0|10),matches([0-9])"`
}