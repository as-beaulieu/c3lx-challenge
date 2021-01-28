package server

import (
	"c3lx-challenge/service"
	"net/http"
)

func heartbeatHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		heartbeatResponse, err := svc.Heartbeat()
		if err != nil {
			respondWithJSON(w, r, 500, r.Body)
			return
		}
		respondWithJSON(w, r, http.StatusOK, heartbeatResponse)
	}
}
