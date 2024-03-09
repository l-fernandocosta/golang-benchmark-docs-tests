// this example of code implements an simple
// webservice
package main

import (
	"log/slog"
	"net/http"

	"github.com/l-fernandodev/test-doc-bench/handlers"
)

const (
	PORT = ":4000"
)

func main() {
	handlers.Routes()

	slog.Info("Listener: Started :  Listening on \":4000\"")
	http.ListenAndServe(PORT, nil)
}
