package delivery

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"service/internal/usecase"
	"service/models"
	"strconv"
	"testing"
)

type fields struct {
	Usecase usecase.IUsecase
}

type args struct {
	r            *http.Request
	response     http.Response
	statusReturn models.StatusCode
}

var testPersons []models.Person = []models.Person{
	models.Person{
		ID:      1,
		Name:    "Phil",
		Address: "strasse der einheit 1",
		Age:     18,
		Work:    "Programmer",
	},
	models.Person{
		ID:      2,
		Name:    "Bob",
		Address: "South Butovo",
		Age:     25,
		Work:    "Nowhere",
	},
}

func TestPersonHandler_GetPerson(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := usecase.NewMockIUsecase(ctl)

	tests := []struct {
		person    models.Person
		body      []byte
		returnErr *models.ErrorResponse
		fields    fields
		args      args
	}{
		{
			person: testPersons[0],
			fields: fields{Usecase: mockUsecase},
			args: args{
				r:            httptest.NewRequest("Get", "http://example.com/api/v1/persons/1", nil),
				statusReturn: models.Okay,
				response:     http.Response{StatusCode: http.StatusOK},
			},
			returnErr: nil,
		},
		{
			person: testPersons[1],
			fields: fields{Usecase: mockUsecase},
			args: args{
				r:            httptest.NewRequest("Get", "http://example.com/api/v1/persons/2", nil),
				statusReturn: models.NotFound,
				response:     http.Response{StatusCode: http.StatusNotFound},
			},
			returnErr: &models.ErrorResponse{Message: "Not found person for id"},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.Okay {
			mockUsecase.EXPECT().GetPerson(models.Person{ID: tests[i].person.ID}).Return(tests[i].person, models.Okay)
			continue
		}
		mockUsecase.EXPECT().GetPerson(models.Person{ID: tests[i].person.ID}).Return(tests[i].person, models.NotFound)
	}

	for _, tt := range tests {
		t.Run(tt.person.Name, func(t *testing.T) {
			tt.args.r = mux.SetURLVars(tt.args.r, map[string]string{"id": strconv.Itoa(tt.person.ID)})
			h := &PersonHandler{
				usecase: tt.fields.Usecase,
			}

			w := httptest.NewRecorder()
			h.GetPerson(w, tt.args.r)
			if tt.args.response.StatusCode != w.Code {
				t.Error(tt.person.Name)
			}
		})
	}
}

func TestPersonHandler_GetPersonsList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := usecase.NewMockIUsecase(ctl)

	tests := []struct {
		persons []models.Person
		fields  fields
		args    args
	}{
		{
			persons: []models.Person{testPersons[0], testPersons[1]},
			fields:  fields{Usecase: mockUsecase},
			args: args{
				r:            httptest.NewRequest("Get", "http://example.com/api/v1/persons", nil),
				statusReturn: models.Okay,
				response:     http.Response{StatusCode: http.StatusOK},
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.Okay {
			mockUsecase.EXPECT().GetPersonsList().Return(tests[i].persons, models.Okay)
			continue
		}
	}

	for _, tt := range tests {
		t.Run("List", func(t *testing.T) {
			h := &PersonHandler{
				usecase: tt.fields.Usecase,
			}

			w := httptest.NewRecorder()
			h.GetPersonsList(w, tt.args.r)
			if tt.args.response.StatusCode != w.Code {
				t.Error("Some error in lest fetching")
			}
		})
	}
}

func TestPersonHandler_AddPerson(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUsecase := usecase.NewMockIUsecase(ctl)

	testBodyRequest1, _ := json.Marshal(models.Person{
		Name:    testPersons[0].Name,
		Work:    testPersons[0].Work,
		Address: testPersons[0].Address,
		Age:     testPersons[0].Age,
	})

	reqReader1 := bytes.NewReader(testBodyRequest1)

	tests := []struct {
		person    models.Person
		body      []byte
		returnErr *models.ValidationErrorResponse
		fields    fields
		args      args
	}{
		{
			person: testPersons[0],
			fields: fields{Usecase: mockUsecase},
			args: args{
				r:            httptest.NewRequest("POST", "http://example.com/api/v1/persons", reqReader1),
				statusReturn: models.Okay,
				response:     http.Response{StatusCode: http.StatusOK},
			},
			returnErr: nil,
		},
	}

	for i := 0; i < len(tests); i++ {
		if tests[i].args.statusReturn == models.Okay {
			mockUsecase.EXPECT().CreatePerson(models.Person{
				Name:    testPersons[0].Name,
				Work:    testPersons[0].Work,
				Address: testPersons[0].Address,
				Age:     testPersons[0].Age,
			}).Return(tests[i].person, models.Created)
			continue
		}
	}

	for _, tt := range tests {
		t.Run(tt.person.Name, func(t *testing.T) {
			h := &PersonHandler{
				usecase: tt.fields.Usecase,
			}

			w := httptest.NewRecorder()
			h.AddPerson(w, tt.args.r)
			if tt.args.response.StatusCode != w.Code {
				t.Error(tt.person.Name)
			}
		})
	}
}

func TestPersonHandler_RemovePerson(t *testing.T) {

}

func TestPersonHandler_UpdatePerson(t *testing.T) {

}
