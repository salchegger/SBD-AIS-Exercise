package httptools

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var BadUrlParamError = errors.New("bad url param")

// ParseIntUrlParam parses URL param with paramName into int, i.e. /api/{someInt}
func ParseIntUrlParam(paramName string, r *http.Request) (int, error) {
	id := chi.URLParam(r, paramName)
	if id == "" {
		return -1, BadUrlParamError
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return -1, BadUrlParamError
	}
	return intId, nil
}

// ParseUintUrlParam parses URL param with paramName into uint, i.e. i.e. /api/{someUint}
func ParseUintUrlParam(paramName string, r *http.Request) (uint, error) {
	i, err := ParseIntUrlParam(paramName, r)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}
