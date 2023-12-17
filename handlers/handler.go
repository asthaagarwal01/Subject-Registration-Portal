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

func (h *RegistrationHandler) Create(ctx *gofr.Context) {
    var registration models.Registration
    if err := ctx.Bind(&registration); err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        return
    }

    err := h.store.Create(&registration)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err.Error())
        return
    }

    ctx.JSON(http.StatusCreated, registration)
}


