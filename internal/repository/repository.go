package repository

import "fmt"

type Repository struct {
	Db Postgres
}

func Init(connStr string) *Repository {

	var postgresDB Postgres

	postgresDB.InitDB(connStr)

	return &Repository{Db: postgresDB}
}

func (r *Repository) RegistrateUser(name, surname, login, password string) error {
	err := r.Db.InsertNewUser(name, surname, login, password)

	return err
}

func (r *Repository) CheckUser(login, password string) (map[string]string, error) {
	user, err := r.Db.GetUser(login)

	if err != nil {
		return nil, err
	}

	if user["password"] == password {
		return user, nil
	}

	return nil, fmt.Errorf("no user!")
}
