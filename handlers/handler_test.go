// Example that shows how to test an
// execution of an endpoint
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/l-fernandodev/test-doc-bench/handlers"
)

const (
	checkmark = "\u2713"
	ballotx   = "\u2717"
)

func init() {
	handlers.Routes()
}

// TestSendJSON testing internal endpoint
func TestSendJSON(t *testing.T) {

	t.Log("Given the need to test SendJSON endpoint")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)

		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotx, err)
		}

		t.Log("\tShould be able to create a request", checkmark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != http.StatusOK {
			t.Fatal("\tShould receive \"200\"", ballotx, rw.Code)
		}

		t.Log("\tShould receive \"200\"", checkmark)

		u := struct {
			Name, Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotx)
		}

		t.Log("\tShould decode the response", checkmark)

		if u.Name == "Fernando" {
			t.Log("\tShould have a Name", checkmark)
		} else {
			t.Error("\tShould have a Name", ballotx)
		}

		if u.Email == "fernandocostadev98@gmail.com" {
			t.Log("\tShould ahve a Email", checkmark)
		} else {
			t.Error("\tShould ahve a email", ballotx)
		}
	}
}
