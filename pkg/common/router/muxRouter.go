package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// region Struct

// MuxRouter contains the required params for router
type MuxRouter struct {
	// instance contains the mux router instance
	instance *mux.Router
}

//endregion

// region Ctor

// NewMuxRouter - it returns new MuxRouter instance
func NewMuxRouter() *MuxRouter {
	return &MuxRouter{instance: mux.NewRouter()}
}

//endregion

// region public methods

// POST - it registers a new POST route
func (mr *MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("POST")
}

// GET - it registers a new GET route
func (mr *MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("GET")
}

// PUT - it registers a new PUT route
func (mr *MuxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("PUT")
}

// DELETE - it registers a new DELETE route
func (mr *MuxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("DELETE")
}

// Serve - it listens and serve the requests
func (mr *MuxRouter) Serve(api string, port string) error {
	fmt.Printf("%v running on port %v", api, port)
	return http.ListenAndServe(port, mr.instance)
}

//endregion
