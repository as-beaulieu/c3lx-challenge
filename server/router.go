package server

import (
	"c3lx-challenge/service"
	"github.com/gorilla/mux"
	"net/http"
)

func makeRouter(svc service.Service) http.Handler {
	muxRouter := mux.NewRouter()

	//Heartbeat and maintenance functions
	muxRouter.HandleFunc("/heartbeat", heartbeatHandler(svc)).Methods("GET")

	//muxRouter.HandleFunc("/one", oneHandler(svc)).Methods("GET")

	return muxRouter
}
