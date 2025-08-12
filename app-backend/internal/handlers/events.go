package handlers

import (
	stderrors "errors"
	"log"
	"net/http"
	"strconv"
	"willdo/internal/errors"
	"willdo/internal/models"
	"willdo/internal/repository"
	"willdo/internal/utils"

	"github.com/gorilla/mux"
)

// EventHandler handles HTTP requests for events
type EventHandler struct {
	logger *log.Logger
	repo   repository.EventRepository
}

func NewEventHandler(logger *log.Logger, repo repository.EventRepository) *EventHandler {
	return &EventHandler{
		logger: logger,
		repo:   repo,
	}
}

func getEventIDFromURL(r *http.Request) int {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// Should never happen as the router ensures that this is a valid number
		panic(err)
	}
	return id
}

// ListAll handles GET requests by returning a list of all events from the database
func (h *EventHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("[DEBUG] get all events")
	rw.Header().Add("Content-Type", "application/json")

	events, err := h.repo.GetAll()

	switch err {
	case nil:
		// Event found
	case errors.ErrEventNotFound:
		h.logger.Println("[ERROR] fetching all events", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&errors.GenericError{Message: "Failed to fetch all events"}, rw)
		return
	default:
		h.logger.Println("[ERROR] fetching all events", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&errors.GenericError{Message: "Failed to fetch all events"}, rw)
		return
	}

	err = utils.ToJSON(events, rw)
	if err != nil {
		h.logger.Println("[ERROR] serializing event", err)
	}
}

// ListSingleEvent handles GET requests by returning a single specified event from the database
func (h *EventHandler) ListSingleEvent(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getEventIDFromURL(r)

	h.logger.Println("[DEBUG] get record id", id)

	event, err := h.repo.GetByID(id)

	switch err {
	case nil:
		// Event found
	case errors.ErrEventNotFound:
		h.logger.Println("[ERROR] fetching event", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&errors.GenericError{Message: err.Error()}, rw)
		return
	default:
		h.logger.Println("[ERROR] fetching event", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&errors.GenericError{Message: err.Error()}, rw)
		return
	}

	err = utils.ToJSON(event, rw)
	if err != nil {
		h.logger.Println("[ERROR] serializing event", err)
	}
}

// Create handles POST requests by creating a new event in the database
func (h *EventHandler) Create(rw http.ResponseWriter, r *http.Request) {
	event := r.Context().Value(EventKey{}).(*models.Event)
	h.logger.Printf("[DEBUG] Creating event: %#v", event)
	err := h.repo.Create(event)
	if err != nil {
		h.logger.Println("[ERROR] creating event", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&errors.GenericError{Message: "Failed to create an event"}, rw)
		return
	}
}

// Update handles PUT requests by updating an existing event
func (h *EventHandler) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	updatedEvent := r.Context().Value(EventKey{}).(*models.Event)
	id := getEventIDFromURL(r)
	h.logger.Println("[DEBUG] updating record with id", id)

	existingEvent, err := h.repo.GetByID(id)
	if stderrors.Is(err, errors.ErrEventNotFound) {
		h.logger.Println("[ERROR] event not found: ", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&errors.GenericError{Message: "Event not found in database"}, rw)
		return
	}

	// Updating only the fields that were provided in the request
	// Keeping the existing values for the rest
	existingEvent.ID = id
	if updatedEvent.Description != "" {
		existingEvent.Description = updatedEvent.Description
	}
	if updatedEvent.DueDate != "" {
		existingEvent.DueDate = updatedEvent.DueDate
	}

	existingEvent.Completed = updatedEvent.Completed

	err = h.repo.Update(existingEvent)
	if stderrors.Is(err, errors.ErrEventNotFound) {
		h.logger.Println("[ERROR] event not found: ", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&errors.GenericError{Message: "Event not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Delete handles DELETE requests by removing an event from the database
func (h *EventHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	id := getEventIDFromURL(r)
	h.logger.Println("[DEBUG] deleting record with id", id)

	err := h.repo.Delete(id)
	if stderrors.Is(err, errors.ErrEventNotFound) {
		h.logger.Println("[ERROR] event not found: ", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&errors.GenericError{Message: "Event not found in database"}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
