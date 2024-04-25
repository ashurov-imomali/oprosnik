package models

import "github.com/lib/pq"

type User struct {
	Id       int64
	Username string
	Password string
	Role     string
	FullName string
}

type Questions struct {
	Id          int64
	Name        string
	Description string
	Variants    pq.StringArray `gorm:"column:variants;type:text[]"`
}

func (Questions) TableName() string {
	return "questions"
}

type Answers struct {
	UserId     int64
	QuestionId int64
	Variant    string
}
