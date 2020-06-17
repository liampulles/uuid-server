package api

import (
	"fmt"
	"net/http"

	"github.com/liampulles/uuid-server/pkg/uuid"

	"github.com/liampulles/uuid-server/pkg/logger"
)

// UUIDHandler implements the http.Handler interface
type UUIDHandler struct {
	loggerService logger.Service
	uuidService   uuid.Service
}

// Check we implement the interface
var _ http.Handler = &UUIDHandler{}

// NewUUIDHandler is a constructor
func NewUUIDHandler(loggerService logger.Service, uuidService uuid.Service) *UUIDHandler {
	return &UUIDHandler{
		loggerService: loggerService,
		uuidService:   uuidService,
	}
}

// ServeHTTP implements the http.Handler interface
func (uh *UUIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gen, err := uh.uuidService.GenerateVersion4UUID()
	if err != nil {
		w.WriteHeader(500)
		if _, errResp := fmt.Fprintf(w, "Encountered an error: %s", err.Error()); errResp != nil {
			uh.loggerService.Errorf("Could not write error to response: %v", errResp)
		}
		uh.loggerService.Errorf("UUID Service error: %v", err)
		return
	}
	w.WriteHeader(200)
	if _, err := fmt.Fprint(w, gen); err != nil {
		uh.loggerService.Errorf("Could not write response: %v", err)
	}
	uh.loggerService.Infof("Generated: %s", gen)
}
