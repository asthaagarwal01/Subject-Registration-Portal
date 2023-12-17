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
    var detail models.RegistrationDetail
    if err := ctx.Bind(&detail); err != nil {
        return nil, NewHTTPError(http.StatusBadRequest, err.Error())
    }

    err := h.store.Create(&detail)  
    if err != nil {
        return nil, NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return detail, nil
}


func (h *RegistrationHandler) GetAll(ctx *gofr.Context) (interface{}, error) {
    registrations, err := h.store.GetAll()
    if err != nil {
        return nil, NewHTTPError(http.StatusInternalServerError, "Error fetching registrations")
    }
    return registrations, nil
}

func (h *RegistrationHandler) Update(ctx *gofr.Context) (interface{}, error) {
    var updatedDetail models.RegistrationDetail
    if err := ctx.Bind(&updatedDetail); err != nil {
        return nil, NewHTTPError(http.StatusBadRequest, err.Error())
    }

    if err := h.store.Update(&updatedDetail); err != nil {
        return nil, NewHTTPError(http.StatusInternalServerError, err.Error())
    }

    return updatedDetail, nil
}
