package storage

import "DB_project/internal/structers"

func InsertNewMachine(newmachine structers.Machine) error {
	_, err := db.Exec(` INSERT INTO car (name, is_new, brand_id, equipment_id) VALUES ($1, $2, $3, $4)`,
		newmachine.Name, newmachine.IsNew, newmachine.BrandId, newmachine.EquipmentId)
	return err
}

func SelectAllMachines() ([]structers.Machine, error) {
	rows, err := db.Query(`  SELECT * FROM car order by id`)
	if err != nil {
		return nil, err
	}
	var machines []structers.Machine
	for rows.Next() {
		var tmp structers.Machine
		err = rows.Scan(&tmp.Id, &tmp.Name, &tmp.IsNew, &tmp.BrandId, &tmp.EquipmentId)
		if err != nil {
			return nil, err
		}
		machines = append(machines, tmp)
	}
	if err = rows.Close(); err != nil {
		return nil, err
	}
	return machines, nil
}

func DeleteMachine(id int) error {
	_, err := db.Exec(`DELETE from car where(id = $1)`, id)
	return err
}
