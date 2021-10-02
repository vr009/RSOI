package delivery

import (
	"net/http"
	"service/models"
)

func Response(w http.ResponseWriter, code models.StatusCode, description string, body []byte) {
	switch code {
	case models.Okay:
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Description", description)
		return
	case models.Created:
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Description", description)
		return
	case models.NotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Description", description)
		return
	case models.BadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Description", description)
		return
	case models.NoContent:
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Description", description)
		return
	case models.InternalError:
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Description", description)
		return
	}
	if body != nil {
		w.Write(body)
		return
	}
}
