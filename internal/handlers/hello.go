package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello ...
type Hello struct {
	logger *log.Logger
}

// NewHello ...
func NewHello(l *log.Logger) *Hello {
	return &Hello{
		logger: l,
	}
}

// Hello ...
func (h *Hello) Hello(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("home route called")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {

		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)

}
