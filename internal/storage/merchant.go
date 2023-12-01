package storage

import "DB_project/internal/structers"

func InsertNewMerchant(newmerchant structers.Merchant) error {
	_, err := db.Exec(`
	INSERT INTO merchant
		(name, country, city, address)
	VALUES ($1, $2, $3, $4)`,
		newmerchant.Name, newmerchant.Country, newmerchant.City, newmerchant.Address)
	return err
}

func SelectAllMerchants() ([]structers.Merchant, error) {
	rows, err := db.Query(`
	SELECT * from merchant order by id 
`)
	if err != nil {
		return nil, err
	}

	var merchants []structers.Merchant
	for rows.Next() {
		var tmp structers.Merchant
		err = rows.Scan(&tmp.Id, &tmp.Name, &tmp.Country, &tmp.City, &tmp.Address)
		if err != nil {
			return nil, err
		}
		merchants = append(merchants, tmp)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return merchants, nil
}

func DeleteMerchant(id int) error {
	_, err := db.Exec(`
	DELETE from merchant
	WHERE (id = $1)`, id)
	return err
}
