package handler

import (
	"cursor-response/model"
	"cursor-response/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *service.PokemonService
}

func NewHandler(service *service.PokemonService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetPokemons(w http.ResponseWriter, r *http.Request) {
	cursorStr := r.URL.Query().Get("cursor")
	var cursor *model.Cursor

	if cursorStr != "" {
		decoded, err := model.DecodeCursor(cursorStr)
		if err != nil {
			http.Error(w, "Invalid cursor", http.StatusBadRequest)
			return
		}
		cursor = decoded
	}

	pokemons, nextCursor, err := h.Service.FetchPokemons(cursor, 10)
	if err != nil {
		http.Error(w, "failed to fetch pokemons", http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"results": pokemons,
	}

	if nextCursor != nil {
		cursorStr, _ := nextCursor.Encode()
		response["next_cursor"] = cursorStr
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
