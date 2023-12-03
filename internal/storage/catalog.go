package storage

import (
	"DB_project/internal/structers"
)

func SelectCatalog() ([]structers.Car, error) {
	rows, err := db.Query(`
	SELECT Catalog.id, Car.name, Car.is_new, Supplier.brand, Equipment.engine, Equipment.color, Equipment.transmission, Equipment.body, Catalog.price
	FROM Car
         	INNER JOIN Supplier ON Car.brand_id = Supplier.id
         	INNER JOIN Equipment ON Car.equipment_id = Equipment.id
         	INNER JOIN Catalog ON Catalog.product_id = Car.id
	ORDER BY Catalog.id ASC;
	`)
	if err != nil {
		return nil, err
	}
	var cars []structers.Car
	for rows.Next() {
		var tmp structers.Car
		err = rows.Scan(&tmp.Id, &tmp.Name, &tmp.IsNew, &tmp.Brand,
			&tmp.Engine, &tmp.Color, &tmp.Transmission, &tmp.Body, &tmp.Price)
		if err != nil {
			return nil, err
		}
		cars = append(cars, tmp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return cars, nil
}

func InsertNewCatalog(newcatalog structers.Catalog) error {
	_, err := db.Exec(`INSERT INTO catalog (id, merchant_id, price, sale, product_id) VALUES ($1, $2, $3, $4)`,
		newcatalog.MerchantId, newcatalog.Price, newcatalog.Sale, newcatalog.ProductId)
	return err
}

func UpdatePriceSale(id, price, sale int) error {
	_, err := db.Exec(`UPDATE catalog SET price = $1, sale = $2 WHERE id = $3`,
		price, sale, id)
	return err
}

func DeleteCatalog(id int) error {
	_, err := db.Exec(`DELETE from catalog where(id = $1)`, id)
	return err
}
