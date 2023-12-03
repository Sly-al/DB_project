package storage

import (
	"DB_project/internal/structers"
	"time"
)

func CreateNewOrder(catalog_pos, client_id int, delivery time.Time) error {
	_, err := db.Exec(`INSERT INTO "Order" (catalog_pos, created_at, delivered_at, client_id) VALUES ($1, $2, $3, $4)`,
		catalog_pos, time.Now(), delivery, client_id)
	return err
}

func SelectAllOrders(clientid int) ([]structers.Car, error) {
	rows, err := db.Query(`
	SELECT Catalog.id, Car.name, Car.is_new, Supplier.brand, Equipment.engine, Equipment.color, Equipment.transmission, Equipment.body, Catalog.price
	FROM Catalog
	INNER JOIN Car ON Catalog.product_id = Car.id
	INNER JOIN Equipment ON Car.equipment_id = Equipment.id
	INNER JOIN Supplier ON Car.brand_id = Supplier.id
	INNER JOIN "Order" ON Catalog.id = "Order".catalog_pos
	WHERE "Order".client_id = $1
	order by Catalog.id;
	`, clientid)
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
