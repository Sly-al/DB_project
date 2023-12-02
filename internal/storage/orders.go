package storage

import "time"

func CreateNewOrder(catalog_pos int, delivery time.Time) error {
	ClientId = 1
	_, err := db.Exec(`INSERT INTO "Order" (catalog_pos, created_at, delivered_at, client_id) VALUES ($1, $2, $3, $4)`,
		catalog_pos, time.Now(), delivery, ClientId)
	return err
}
