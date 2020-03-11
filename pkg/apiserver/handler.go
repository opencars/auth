package apiserver

import (
	"encoding/json"
	"log"
	"net/http"
)

// The Handler helps to handle errors in one place.
type Handler func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		switch e := err.(type) {
		case Error:
			w.WriteHeader(e.Code)
			if err := json.NewEncoder(w).Encode(e); err != nil {
				panic(err)
			}
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			e = NewError(http.StatusInternalServerError, "system.unhealthy")
			if err := json.NewEncoder(w).Encode(e); err != nil {
				panic(err)
			}
		}
	}
}
