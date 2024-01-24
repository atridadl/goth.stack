package api

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func Ping(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong!"))
	return nil
}
