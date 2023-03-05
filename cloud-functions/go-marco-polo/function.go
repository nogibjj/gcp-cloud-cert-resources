// Package p contains an HTTP Cloud Function.
package p

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// make this into a Marco Polo function that accepts name: "Marco" and returns "Polo"
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		switch err {
		case io.EOF:
			fmt.Fprint(w, "Send me a JSON object with a name property.")
			return
		default:
			log.Printf("json.NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	// if name isn't Marco send back I need another name
	if d.Name == "Marco" {
		fmt.Fprint(w, "Polo!")
		return
	} else {
		fmt.Fprint(w, "I need another name")
		return
	}
}
