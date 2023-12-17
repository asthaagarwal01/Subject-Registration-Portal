package configs

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

func NewDB() (*gorm.DB, error) {
    
    dsn := "root:password@tcp(localhost:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
