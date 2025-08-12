package handlers

import (
	"context"
	"log"
	"net/http"
	"willdo/internal/errors"
	"willdo/internal/models"
	"willdo/internal/utils"
	"willdo/internal/validator"
)

// EventKey is used for storing Event in the request context
type EventKey struct{}

type EventValidationMiddleware struct {
	logger    *log.Logger
	validator *validator.Validator
}

func NewEventValidationMiddleware(logger *log.Logger, validator *validator.Validator) *EventValidationMiddleware {
	return &EventValidationMiddleware{
		logger:    logger,
		validator: validator,
	}
}

func (m *EventValidationMiddleware) ValidateEvent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		e := &models.Event{}

		err := utils.FromJSON(e, r.Body)
		if err != nil {
			m.logger.Println("[ERROR] deserializing event", err)
			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&errors.GenericError{Message: err.Error()}, rw)
			return
		}

		// Validating the event
		errs := m.validator.Validate(e)
		if len(errs) != 0 {
			m.logger.Println("[ERROR] validating event", errs)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			utils.ToJSON(&errors.ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), EventKey{}, e)
		r = r.WithContext(ctx)

		// Calling next HTTP handler
		next.ServeHTTP(rw, r)
	})
}
