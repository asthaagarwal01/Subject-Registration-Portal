package models

type RegistrationDetail struct {
	Name        string `json:"name" gorm:"column:name;type:varchar(30);not null"`
	RollNo      int    `json:"rollNo" gorm:"column:roll_no;type:int;primaryKey"`
	SubjectName string `json:"subjectName" gorm:"column:subject_name;type:varchar(20);not null"`
	Status      string `json:"status" gorm:"column:status;type:varchar(20);not null;default:'pending'"`
}
