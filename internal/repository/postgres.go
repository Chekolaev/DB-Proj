package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) InitDB(connStr string) {
	d, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalln("db init failed:", err)
	} else if err := d.Ping(); err != nil {
		log.Fatalln("db connect failed:", err)
	}

	p.db = d
}

func (p *Postgres) InsertNewUser(name, surname, login, password string) error {
	_, err := p.db.Exec("INSERT INTO users (name, surname, login, password, role) VALUES($1, $2, $3, $4, 1)", name, surname, login, password)

	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}

	return nil
}

func (p *Postgres) GetUser(login string) (map[string]string, error) {
	rows, err := p.db.Query("SELECT uuid, name, surname, login, password, role FROM users WHERE login = $1", login)

	if err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}

	User := struct {
		uuid     string
		name     string
		surname  string
		login    string
		password string
		role     string
	}{}

	if rows.Next() {
		err := rows.Scan(&User.uuid, &User.name, &User.surname, &User.login, &User.password, &User.role)
		if err != nil {
			return nil, fmt.Errorf("saving user info failed: %w", err)
		}

		UserData := map[string]string{
			"uuid":     User.uuid,
			"name":     User.name,
			"surname":  User.surname,
			"login":    User.login,
			"password": User.password,
			"role":     User.role,
		}
		return UserData, nil
	}

	return nil, fmt.Errorf("no user!")
}

func (p *Postgres) GetUserByInterface(uuid interface{}) (map[string]string, error) {
	rows, err := p.db.Query("SELECT uuid, name, surname, login, password, role FROM users WHERE uuid = $1", uuid)

	if err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}
	defer rows.Close()

	User := struct {
		uuid     string
		name     string
		surname  string
		login    string
		password string
		role     string
	}{}

	if rows.Next() {
		err := rows.Scan(&User.uuid, &User.name, &User.surname, &User.login, &User.password, &User.role)
		if err != nil {
			return nil, fmt.Errorf("saving user info failed: %w", err)
		}

		UserData := map[string]string{
			"uuid":     User.uuid,
			"name":     User.name,
			"surname":  User.surname,
			"login":    User.login,
			"password": User.password,
			"role":     User.role,
		}
		return UserData, nil
	}

	return nil, fmt.Errorf("no user!")
}

func (p *Postgres) GetAllBooks() ([]map[string]string, error) {
	rows, err := p.db.Query("SELECT uuid, name, description FROM books WHERE status = 1")

	if err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}

	defer rows.Close()

	Book := struct {
		uuid        string
		name        string
		description string
	}{}

	var resp []map[string]string

	for rows.Next() {
		err := rows.Scan(&Book.uuid, &Book.name, &Book.description)

		if err != nil {
			return nil, fmt.Errorf("book scan failed: %w", err)
		}

		resp = append(resp, map[string]string{
			"uuid":        Book.uuid,
			"name":        Book.name,
			"description": Book.description,
		})
	}

	return resp, nil
}

func (p *Postgres) GetBookByUUID(uuidBook string) (map[string]string, error) {
	rows, err := p.db.Query("SELECT uuid, name, description FROM books WHERE status = 1 AND uuid = $1", uuidBook)

	if err != nil {
		return nil, fmt.Errorf("db query failed: %w", err)
	}

	defer rows.Close()

	Book := struct {
		uuid        string
		name        string
		description string
	}{}

	if rows.Next() {
		err := rows.Scan(&Book.uuid, &Book.name, &Book.description)

		if err != nil {
			return nil, fmt.Errorf("book scan failed: %w", err)
		}

		resp := map[string]string{
			"uuid":        Book.uuid,
			"name":        Book.name,
			"description": Book.description,
		}

		return resp, nil
	}
	return nil, nil
}

func (p *Postgres) SetNewBookStatus(uuidBook, statusNum string, UUID interface{}) (string, error) {
	if statusNum == "3" {
		_, err := p.db.Exec("UPDATE books SET status = $1 WHERE uuid = $2", statusNum, uuidBook)

		if err != nil {
			return "", fmt.Errorf("update db failed: %w", err)
		}
	} else if statusNum == "2" {
		_, err := p.db.Exec("UPDATE books SET status = $1, holder_uuid = $2 WHERE uuid = $3", statusNum, UUID, uuidBook)

		if err != nil {
			return "", fmt.Errorf("update db failed: %w", err)
		}
	} else if statusNum == "1" {
		_, err := p.db.Exec("UPDATE books SET status = $1, holder_uuid = NULL WHERE uuid = $2", statusNum, uuidBook)

		if err != nil {
			return "", fmt.Errorf("update db failed: %w", err)
		}
	}

	return "Change success!", nil
}

func (p *Postgres) ShowRequests() ([]map[string]string, error) {
	rows, err := p.db.Query("SELECT b.name, u.login FROM books AS b INNER JOIN users AS u ON u.uuid = b.holder_uuid WHERE b.status = 2")

	if err != nil {
		return nil, fmt.Errorf("db query error in ShowRequests: %w", err)
	}

	var resp []map[string]string

	req := struct {
		bookName  string
		userLogin string
	}{}

	for rows.Next() {
		err := rows.Scan(&req.bookName, &req.userLogin)

		if err != nil {
			return nil, fmt.Errorf("db exec failed: %w", err)
		}

		resp = append(resp, map[string]string{
			"book_name":  req.bookName,
			"user_login": req.userLogin,
		})
	}

	return resp, nil
}
