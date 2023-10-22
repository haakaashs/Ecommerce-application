package models

type Student struct {
	Id    uint
	Name  string
	Score uint
}

func (*Student) TableName() string {
	return "students"
}
