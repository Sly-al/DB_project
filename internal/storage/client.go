package storage

import (
	"DB_project/internal/structers"
)

var ClientId int

func UpdateStatus(id int, status string) error {
	_, err := db.Exec(`
	UPDATE client SET status = $1 WHERE status = $2`, status, id)
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
