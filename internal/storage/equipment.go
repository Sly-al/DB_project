package storage

import "DB_project/internal/structers"

func InsertNewEquipment(newequipment structers.Equipment) error {
	_, err := db.Exec(`INSERT INTO equipment (engine, color, transmission, body)
VALUES ($1, $2, $3, $4)`,
		newequipment.Engine, newequipment.Color, newequipment.Transmission, newequipment.Body)
	return err
}

func SelectAllEquipment() ([]structers.Equipment, error) {
	rows, err := db.Query(`  SELECT * FROM equipment order by id`)
	if err != nil {
		return nil, err
	}
	var equipment []structers.Equipment
	for rows.Next() {
		var tmp structers.Equipment
		err = rows.Scan(&tmp.Id, &tmp.Engine, &tmp.Color, &tmp.Transmission, &tmp.Body)
		if err != nil {
			return nil, err
		}
		equipment = append(equipment, tmp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return equipment, nil
}

func DeleteEquipment(id int) error {
	_, err := db.Exec(`DELETE from equipment where(id = $1)`, id)
	return err
}
