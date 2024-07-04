package entities

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionText  string    `gorm:"type:varchar(500);not null"`
    OptionA       string    `gorm:"type:varchar(200);not null"`
    OptionB       string    `gorm:"type:varchar(200);not null"`
    OptionC       string    `gorm:"type:varchar(200);not null"`
    OptionD       string    `gorm:"type:varchar(200);not null"`
    CorrectOption string    `gorm:"type:char(1);not null"`
}
