package docs

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name, Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		slog.Error(err.Error())
	}

	// use fmt to write at `stdout`
	fmt.Println(u)
	// Output:
	// { Fernando fernandocostadev98@gmail.com }
}
