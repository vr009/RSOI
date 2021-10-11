package delivery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"service/models"
	"testing"
)

var codes = []models.StatusCode{
	models.Okay,
	models.NotFound,
	models.BadRequest,
	models.Created,
	models.NoContent,
	models.InternalError,
}

var results = []int{
	http.StatusOK,
	http.StatusNotFound,
	http.StatusBadRequest,
	http.StatusCreated,
	http.StatusNoContent,
	http.StatusInternalServerError,
}

func TestResponse(t *testing.T) {
	for _, tt := range codes {
		w := httptest.NewRecorder()
		Response(w, codes[tt], fmt.Sprintf("test %d", codes[tt]), nil)
		if w.Code != results[tt] {
			t.Error("Incorrect responser\n")
		}
	}
}
