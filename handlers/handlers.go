// handlers package expose the endpoints
// to the webservice
package handlers

import (
	"encoding/json"
	"net/http"
)

// Routes defines the routes to the webservice
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

// SendJSON returns an simples JSON document
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name, Email string
	}{
		Name:  "Fernando",
		Email: "fernandocostadev98@gmail.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
