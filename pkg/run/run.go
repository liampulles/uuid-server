package run

import (
	"fmt"
	"log"
	"net/http"

	googleUuid "github.com/google/uuid"
	"github.com/liampulles/go-config"

	"github.com/liampulles/uuid-server/pkg/api"
	uuidConfig "github.com/liampulles/uuid-server/pkg/config"
	"github.com/liampulles/uuid-server/pkg/logger"
	"github.com/liampulles/uuid-server/pkg/uuid"
)

// Run will run the uuid-server
func Run(source config.Source, listenAndServe func(string, http.Handler) error) int {
	cfg, err := uuidConfig.InitUUIDServerConfig(source)
	if err != nil {
		fmt.Printf("Error parsing config: %v", err)
		return 1
	}

	logSvc := logger.NewServiceImpl(cfg.LogLevel, log.Printf)
	uuidSvc := uuid.NewServiceImpl(googleUuid.NewRandom)
	handler := api.NewUUIDHandler(logSvc, uuidSvc)

	logSvc.Infof("Now serving on port %d!", cfg.Port)
	if err := listenAndServe(fmt.Sprintf(":%d", cfg.Port), handler); err != nil {
		fmt.Printf("Error during serving: %v", err)
		return 2
	}
	return 0
}
