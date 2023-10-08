package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("Ping handler called")
	writeSuccessResponse(w, struct{ Resp string }{"Pong"}, http.StatusOK)
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}, status int) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("error during response marshal: ", data)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(bytes)
	if err != nil {
		log.Println("error during writing to response writer")
	}
}
