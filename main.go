package main

import (
    "subject-registration/configs"
    "subject-registration/models"
    "subject-registration/store"
    "subject-registration/handlers"
    "gofr.dev/pkg/gofr"
)

func main() {
    app := gofr.New()

    db, err := configs.NewDB()
    if err != nil {
        panic("failed to connect database")
    }

   
    db.AutoMigrate(&models.RegistrationDetail{})

    registrationStore := store.NewRegistrationStore(db)
    registrationHandler := handlers.NewRegistrationHandler(registrationStore)

    app.POST("/registrations", registrationHandler.Create)
    app.GET("/registrations", registrationHandler.GetAll)
    app.PUT("/registrations/{id}", registrationHandler.Update) 
    app.Start()
}
