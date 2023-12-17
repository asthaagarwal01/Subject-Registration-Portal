package store

import (
    "subject-registration/models"
    "gorm.io/gorm"
)

type RegistrationStore interface {
    Create(detail *models.RegistrationDetail) error
    GetAll() ([]models.RegistrationDetail, error)
    Update(detail *models.RegistrationDetail) error
}

type registrationStore struct {
    db *gorm.DB
}



func NewRegistrationStore(db *gorm.DB) RegistrationStore {
    return &registrationStore{db: db}
}

func (store *registrationStore) Create(detail *models.RegistrationDetail) error {
    return store.db.Create(detail).Error
}

func (store *registrationStore) GetAll() ([]models.RegistrationDetail, error) {
    var details []models.RegistrationDetail
    err := store.db.Find(&details).Error
    return details, err
}

func (store *registrationStore) Update(detail *models.RegistrationDetail) error {
   
    return store.db.Model(&models.RegistrationDetail{}).Where("roll_no = ?", detail.RollNo).Updates(detail).Error
}