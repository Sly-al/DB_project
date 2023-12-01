package storage

import (
	"DB_project/internal/structers"
)

func InsertNewSupplier(newsupplier structers.Supplier) error {
	_, err := db.Exec(`
	INSERT INTO supplier (brand, country, city, address) VALUES ($1, $2, $3, $4)`,
		newsupplier.Brand, newsupplier.Country, newsupplier.City, newsupplier.Address)
	return err
}

func SelectAllSuppliers() ([]structers.Supplier, error) {
	rows, err := db.Query(`
	SELECT * FROM supplier order by id`)
	if err != nil {
		return nil, err
	}
	var suppliers []structers.Supplier
	for rows.Next() {
		var tmp structers.Supplier
		err = rows.Scan(&tmp.Id, &tmp.Brand, &tmp.Country, &tmp.City, &tmp.Address)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, tmp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return suppliers, nil
}

func DeleteSupplier(id int) error {
	_, err := db.Exec(`
	DELETE from supplier where(id = $1)`, id)
	return err
}
