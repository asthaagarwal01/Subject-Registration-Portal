package configs

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("registration.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
