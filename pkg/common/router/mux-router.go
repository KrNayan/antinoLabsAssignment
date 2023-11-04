package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MuxRouter struct {
	instance *mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{instance: mux.NewRouter()}
}

func (mr *MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("POST")
}

func (mr *MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("GET")
}

func (mr *MuxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("PUT")
}

func (mr *MuxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	mr.instance.HandleFunc(uri, f).Methods("DELETE")
}

func (mr *MuxRouter) Serve(api string, port string) error {
	fmt.Printf("%v running on port %v", api, port)
	return http.ListenAndServe(port, mr.instance)
}
