package storage

import (
	"DB_project/internal/structers"
)

var ClientId int

func UpdateStatus(login string) error {
	_, err := db.Exec(`
	UPDATE client SET status = 'VIP' WHERE login = $1`, login)
	return err
}

func SelectAllClients() ([]structers.Client, error) {
	rows, err := db.Query(`
	SELECT * from client order by id 
`)
	if err != nil {
		return nil, err
	}

	var clients []structers.Client
	for rows.Next() {
		var tmp structers.Client
		err = rows.Scan(&tmp.Id, &tmp.Login, &tmp.Password, &tmp.Surname, &tmp.Name, &tmp.Status)
		if err != nil {
			return nil, err
		}
		tmp.Password = "*"
		clients = append(clients, tmp)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return clients, nil
}

func SelectClientId(login string) (int, error) {
	rows, err := db.Query(`select id from client where login = $1`, login)
	if err != nil {
		return 0, err
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	if err := rows.Close(); err != nil {
		return 0, err
	}
	return id, nil
}

func SelectAllClientLogins() ([]string, error) {
	rows, err := db.Query(`SELECT login from client order by id`)
	if err != nil {
		return nil, err
	}

	var logins []string
	for rows.Next() {
		var login string
		err = rows.Scan(&login)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return logins, nil
}
func SelectAllRegularClientLogins() ([]string, error) {
	rows, err := db.Query(`SELECT login from client where status = 'Regular' order by id`)
	if err != nil {
		return nil, err
	}

	var logins []string
	for rows.Next() {
		var login string
		err = rows.Scan(&login)
		if err != nil {
			return nil, err
		}
		logins = append(logins, login)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return logins, nil
}

func GetPassword(login string) (string, error) {
	var password string
	rows, err := db.Query(`SELECT id, password FROM client where login = $1`, login)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(&ClientId, &password)
		if err != nil {
			return "", err
		}
	}
	if err := rows.Close(); err != nil {
		return "", err
	}
	return password, nil
}

func InsertNewClient(neweclient structers.Client) error {
	_, err := db.Exec(`INSERT INTO client (login, password, surname, name, status)
VALUES ($1, $2, $3, $4, $5)`,
		neweclient.Login, neweclient.Password, neweclient.Surname, neweclient.Name, neweclient.Status)
	return err
}

func DeleteClient(login string) error {
	_, err := db.Exec(`DELETE FROM client where (login = $1)`, login)
	return err
}
