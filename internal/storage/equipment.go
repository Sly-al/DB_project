package storage

import (
	"DB_project/internal/structers"
	"fmt"
)

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

func InsertToEq(eq structers.Equipment) (int, error) {
	rows, err := db.Query(`SELECT id from equipment 
          where engine = $1 and color = $2 and transmission=$3 and body=$4`,
		eq.Engine, eq.Color, eq.Transmission, eq.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var id int
	for rows.Next() {
		err = rows.Scan(&id)
	}
	if id == 0 {
		sqlStatement := `INSERT INTO equipment (engine, color, transmission, body)
		VALUES ($1, $2, $3, $4) RETURNING id`
		err = db.QueryRow(sqlStatement, eq.Engine, eq.Color, eq.Transmission, eq.Body).Scan(&id)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	}
	fmt.Println(id)
	return id, nil
}

func DeleteEquipment(id int) error {
	_, err := db.Exec(`DELETE from equipment where(id = $1)`, id)
	return err
}
