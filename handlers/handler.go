package handlers

import (
    "net/http"
    "subject-registration/models"
    "subject-registration/store"
    "gofr.dev/pkg/gofr"
)

type RegistrationHandler struct {
    store store.RegistrationStore
}

func NewRegistrationHandler(store store.RegistrationStore) *RegistrationHandler {
    return &RegistrationHandler{store: store}
}

func (h *RegistrationHandler) Create(ctx *gofr.Context) (interface{}, error) {
    var registration models.Registration
    if err := ctx.Bind(&registration); err != nil {
        
        return nil, NewHTTPError(http.StatusBadRequest, err.Error())
    }

    err := h.store.Create(&registration)
    if err != nil {
        
        return nil, NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    
    return registration, nil
}