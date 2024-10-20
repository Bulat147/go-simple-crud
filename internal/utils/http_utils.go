package utils

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func ParseUUIDFromRequestParam(param string, r *http.Request) (uuid.UUID, error) {
	idParam := r.URL.Query().Get(param)
	if idParam == "" {
		return uuid.Nil, fmt.Errorf("param %s is empty", param)
	}
	var id, err = uuid.Parse(idParam)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid %s param", param)
	}
	return id, nil
}
