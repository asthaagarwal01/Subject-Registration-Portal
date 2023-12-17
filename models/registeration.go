package models

type RegistrationDetail struct {
    Name        string `json:"name" gorm:"column:name"`
    RollNo      int    `json:"rollNo" gorm:"column:roll_no;primaryKey"`
    SubjectName string `json:"subjectName" gorm:"column:subject_name"`
    Status      string `json:"status" gorm:"column:status"`
    SRNo        int    `json:"srNo" gorm:"column:SR_No;autoIncrement"`
}
