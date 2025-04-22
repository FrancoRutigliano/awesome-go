package service

import (
	"cursor-response/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func extractIDFromURL(url string) (int, error) {
	parts := strings.Split(strings.Trim(url, "/"), "/")
	return strconv.Atoi(parts[len(parts)-1])
}

type PokeAPIResponse struct {
	Results []Pokemon `json:"results"`
}

type PokemonService struct {
}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (s *PokemonService) FetchPokemons(cursor *model.Cursor, pageSize int) ([]Pokemon, *model.Cursor, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=200")
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var apiResponse PokeAPIResponse

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, nil, err
	}

	filtered := []Pokemon{}
	count := 0
	for _, item := range apiResponse.Results {
		id, err := extractIDFromURL(item.Url)
		if err != nil {
			continue
		}
		if cursor == nil || id > cursor.LastId {
			filtered = append(filtered, item)
			count++
			if count >= pageSize {
				next := &model.Cursor{LastId: id}
				return filtered, next, nil
			}
		}
	}
	return filtered, nil, nil
}
