package v1

import (
	"errors"
	"fmt"
	"net/http"
)

func notFound(w http.ResponseWriter, _ *http.Request) {
	code := http.StatusNotFound
	errorMessage := fmt.Sprintf("%d %s", code, http.StatusText(code))
	errorResponse(w, http.StatusNotFound, errors.New(errorMessage))
}
