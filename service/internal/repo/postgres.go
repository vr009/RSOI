package repo

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"service/models"
)

const (
	INSERTQUERY = "INSERT INTO persons.persons(name, age, work, address) VALUES($1, $2, $3, $4) RETURNING person_id;"
	DELETEQUERY = "DELETE FROM WHERE person_id=$1;"
	UPDATEQUERY = "UPDATE SET WHERE person_id=$1;"
	GETQUERY    = "SELECT FROM WHERE person_id=$1;"
	LISTQUERY   = "SELECT * FROM public.persons;"
)

type PersonRepo struct {
	conn *pgxpool.Pool
}

func NewPersonRepo(conn *pgxpool.Pool) *PersonRepo {
	return &PersonRepo{
		conn: conn,
	}
}

func (pr *PersonRepo) CreatePerson(person models.Person) (models.Person, models.StatusCode) {
	NewPerson := models.Person{Name: person.Name, Address: person.Address, Work: person.Work, Age: person.Age}
	row := pr.conn.QueryRow(context.Background(), INSERTQUERY)
	err := row.Scan(&NewPerson.ID)
	if err != nil {
		return models.Person{}, models.BadRequest
	}
	return NewPerson, models.Okay
}

func (pr *PersonRepo) DeletePerson(person models.Person) models.StatusCode {
	_, err := pr.conn.Exec(context.Background(), DELETEQUERY, person.ID)
	if err != nil {
		return models.InternalError
	}
	return models.Okay
}

func (pr *PersonRepo) UpdatePerson(person models.Person) models.StatusCode {
	_, status := pr.GetPerson(person)
	if status != models.Okay {
		return models.NotFound
	}

	_, err := pr.conn.Exec(context.Background(), UPDATEQUERY, person.ID)
	if err != nil {
		return models.BadRequest
	}
	return models.Okay
}
func (pr *PersonRepo) GetPerson(person models.Person) (models.Person, models.StatusCode) {
	rows := pr.conn.QueryRow(context.Background(), GETQUERY, person.ID)
	err := rows.Scan(&person)
	if err != nil {
		return models.Person{}, models.NotFound
	}
	return models.Person{}, models.Okay
}
func (pr *PersonRepo) GetPersonsList() ([]models.Person, models.StatusCode) {
	rows, err := pr.conn.Query(context.Background(), LISTQUERY)
	if err != nil {
		return nil, models.InternalError
	}
	list := make([]models.Person, 0)
	person := models.Person{}
	for rows.Next() {
		rows.Scan(&person)
		list = append(list, person)
	}
	return list, models.Okay
}
