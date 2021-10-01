package delivery

import "service/models"

func PersonIsValid(person models.Person) map[string]string {
	errs := make(map[string]string)
	if person.Work == "" {
		errs["Work"] = "No work info provided"
	}
	if person.Age == 0 {
		errs["Age"] = "No age info provided"
	}
	if person.Address == "" {
		errs["Address"] = "No address info provided"
	}
	if person.Name == "" {
		errs["Name"] = "No name info provided"
	}

	if len(errs) == 0 {
		return nil
	}
	return errs
}
