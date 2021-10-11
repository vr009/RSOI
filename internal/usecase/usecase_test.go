package usecase

import (
	"github.com/golang/mock/gomock"
	"service/internal/repo"
	"service/models"
	"testing"
)

type fields struct {
	repo repo.IRepo
}

type args struct {
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
	models.Person{
		ID:      3,
		Name:    "Drake",
		Address: "California",
		Age:     37,
		Work:    "Nowhere",
	},
}

var codes = []models.StatusCode{
	models.Okay,
	models.NotFound,
}

func TestPersonUsecase_GetPerson(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := repo.NewMockIRepo(ctl)

	tests := []struct {
		person    models.Person
		returnErr *models.ErrorResponse
		fields    fields
		args      args
	}{
		{
			person: testPersons[0],
			fields: fields{repo: mockRepo},
			args: args{
				statusReturn: models.Okay,
			},
			returnErr: nil,
		},
		{
			person: testPersons[1],
			fields: fields{repo: mockRepo},
			args: args{
				statusReturn: models.NotFound,
			},
			returnErr: &models.ErrorResponse{Message: "Not found person for id"},
		},
	}

	for i := 0; i < len(tests); i++ {
		mockRepo.EXPECT().GetPerson(tests[i].person).Return(tests[i].person, tests[i].args.statusReturn)
	}

	for i, tt := range tests {
		t.Run(tt.person.Name, func(t *testing.T) {
			u := &PersonUsecase{
				repo: tt.fields.repo,
			}

			u.GetPerson(tt.person)
			if tt.args.statusReturn != codes[i] {
				t.Error(tt.person.Name)
			}
		})
	}
}

func TestPersonUsecase_CreatePerson(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := repo.NewMockIRepo(ctl)

	tests := []struct {
		person    models.Person
		returnErr *models.ErrorResponse
		fields    fields
		args      args
	}{
		{
			person: testPersons[0],
			fields: fields{repo: mockRepo},
			args: args{
				statusReturn: models.Okay,
			},
			returnErr: nil,
		},
		{
			person: testPersons[1],
			fields: fields{repo: mockRepo},
			args: args{
				statusReturn: models.NotFound,
			},
			returnErr: &models.ErrorResponse{Message: "Not found person for id"},
		},
	}

	for i := 0; i < len(tests); i++ {
		mockRepo.EXPECT().CreatePerson(tests[i].person).Return(tests[i].person, tests[i].args.statusReturn)
	}

	for i, tt := range tests {
		t.Run(tt.person.Name, func(t *testing.T) {
			u := &PersonUsecase{
				repo: tt.fields.repo,
			}

			u.CreatePerson(tt.person)
			if tt.args.statusReturn != codes[i] {
				t.Error(tt.person.Name)
			}
		})
	}
}
