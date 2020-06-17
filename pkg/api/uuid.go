package api

import (
	"fmt"
	"net/http"

	"github.com/liampulles/uuid-server/pkg/uuid"

	"github.com/liampulles/uuid-server/pkg/logger"
)

type UUIDHandler struct {
	loggerService logger.Service
	uuidService   uuid.Service
}

var _ http.Handler = &UUIDHandler{}

func NewUUIDHandler(loggerService logger.Service, uuidService uuid.Service) *UUIDHandler {
	return &UUIDHandler{
		loggerService: loggerService,
		uuidService:   uuidService,
	}
}

func (uh *UUIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gen, err := uh.uuidService.GenerateVersion4UUID()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Encountered an error: %s", err.Error())
		uh.loggerService.Errorf("%v", err)
		return
	}
	w.WriteHeader(200)
	fmt.Fprint(w, gen)
	uh.loggerService.Infof("Generated %s", gen)
}
