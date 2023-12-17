package models

import "gorm.io/gorm"

type Registration struct {
    gorm.Model
    StudentID  uint
    SubjectID  uint
    Status     string // "requested", "approved", "registered"
}
