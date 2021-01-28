package server

import (
	"c3lx-challenge/service"
	"c3lx-challenge/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func getChallengesHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := svc.GetChallenges()
		if err != nil {
			fmt.Println("error creating simple message", err)
			respondWithJSON(w, r, 500, r.Body)
		}
		respondWithJSON(w, r, http.StatusOK, response)
	}
}

func acceptChallengeHandler(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var challengeRequest models.AcceptChallengeRequest

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&challengeRequest); err != nil {
			respondWithJSON(w, r, http.StatusBadRequest, r.Body)
			return
		}
		defer r.Body.Close()

		heartbeatResponse, err := svc.AcceptChallenge(challengeRequest)
		if err != nil {
			respondWithJSON(w, r, 500, r.Body)
			return
		}
		respondWithJSON(w, r, http.StatusOK, heartbeatResponse)
	}
}
