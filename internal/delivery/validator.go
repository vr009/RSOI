package delivery

import (
	"reflect"
	"service/models"
)

const (
	IvalidType = "invalid type"
)

func PersonIsValid(person models.Person) map[string]string {
	errs := make(map[string]string)
	if reflect.TypeOf(person.Work).String() != "string" && person.Work != "" {
		errs["Work"] = IvalidType
	}
	if reflect.TypeOf(person.Age).String() != "int" && person.Age != 0 {
		errs["Age"] = IvalidType
	}
	if reflect.TypeOf(person.Address).String() != "string" && person.Address != "" {
		errs["Address"] = IvalidType
	}
	if reflect.TypeOf(person.Name).String() != "string" && person.Name != "" {
		errs["Name"] = IvalidType
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
