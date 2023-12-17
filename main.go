package main

import (
    "subject-registration/configs"
    "subject-registration/handlers"
    "subject-registration/models"
    "subject-registration/store"
    "github.com/zopsmart/GoFr"
)

func main() {
    app := GoFr.New()

    db, err := configs.NewDB()
    if err != nil {
        panic("failed to connect database")
    }

    // AutoMigrate the schema
    db.AutoMigrate(&models.Registration{})

    registrationStore := store.NewRegistrationStore(db)
    registrationHandler := handlers.NewRegistrationHandler(registrationStore)

    app.POST("/registrations", registrationHandler.Create)
   

    app.Start()
}
