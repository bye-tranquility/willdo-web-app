package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"willdo/internal/handlers"
	"willdo/internal/repository"
	"willdo/internal/validator"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func New(logger *log.Logger, addr string, eventRepo repository.EventRepository) *Server {
	v := validator.New()

	// middleware setup
	// event validation
	evm := handlers.NewEventValidationMiddleware(logger, v)
	// CORS
	ch := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"http://localhost:3000", "http://localhost", "http://frontend:3000"}),
		gohandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE"}),
		gohandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// handlers setup
	eh := handlers.NewEventHandler(logger, eventRepo)

	// routes setup
	router := setupRoutes(eh, evm, logger)

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      ch(router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		logger:     logger,
	}
}

func setupRoutes(eh *handlers.EventHandler, evm *handlers.EventValidationMiddleware, logger *log.Logger) *mux.Router {
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/events", eh.ListAll)
	getR.HandleFunc("/events/{id:[0-9]+}", eh.ListSingleEvent)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/events", eh.Create)
	postR.Use(evm.ValidateEvent)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/events/{id:[0-9]+}", eh.Update)
	putR.Use(evm.ValidateEvent)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/events/{id:[0-9]+}", eh.Delete)
	return sm
}

func (s *Server) Start() error {
	s.logger.Printf("[INFO] Starting server on %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Println("[INFO] Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}
