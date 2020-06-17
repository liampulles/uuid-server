package main

import (
	"net/http"
	"os"

	"github.com/liampulles/go-config"
	"github.com/liampulles/uuid-server/cmd/uuid-server/run"
)

func main() {
	code := run.Run(config.NewEnvSource(), http.ListenAndServe)
	os.Exit(code)
}
