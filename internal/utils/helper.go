package utils

import (
	"encoding/json"
	"net/http"
)

// Função auxiliar para escrever o JSON de retorno em caso de sucesso na requisição
func WriteJSON(w http.ResponseWriter, status int, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}
