package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Error ErrorResponse	`json:"error"`
}

type ErrorResponse struct {
	Status	int		`json:"status"`
	Message	string	`json:"message"`
}

func Response(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&JsonResponse{Error:ErrorResponse{
		Status:  status,
		Message: message,
	}}); err != nil {
		log.Fatalln(err)
	}
}
