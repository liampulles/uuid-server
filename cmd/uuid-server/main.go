package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const (
	defaultPort  = "8080"
	logLevelInfo = iota
	logLevelError
)

func main() {
	port := configPort()
	handler := &uuidHandler{}
	http.ListenAndServe(":"+port, handler)
}

func configPort() string {
	port, present := os.LookupEnv("PORT")
	if !present {
		return defaultPort
	}
	return port
}

type uuidHandler struct{}

var _ http.Handler = &uuidHandler{}

func (uh *uuidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gen, err := generateVersion4UUID()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Encountered an error: %s", err.Error())
		log.Printf("ERROR: %v", err)
		return
	}
	w.WriteHeader(200)
	fmt.Fprint(w, gen)
}

func generateVersion4UUID() (string, error) {
	gen, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return gen.String(), nil
}
