package repositories

import (
	"database/sql"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
)

type PerangkatDesaRepository struct {
	DB *sql.DB
}

func NewPerangkatDesaRepository() *PerangkatDesaRepository {
	// initialize and configure the UserRepository instance
	return &PerangkatDesaRepository{
		// initialize fields and dependencies
	}
}

func (pr *PerangkatDesaRepository) AddPerangkatDesa(perangkat *models.PerangkatDesa) error {
	stmt, err := pr.DB.Prepare("INSERT INTO perangkat (nama, jabatan) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(perangkat.Nama, perangkat.Jabatan)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PerangkatDesaRepository) UpdatePerangkatDesa(perangkat *models.PerangkatDesa) error {
	stmt, err := pr.DB.Prepare("UPDATE perangkat SET nama=?, jabatan=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(perangkat.Nama, perangkat.Jabatan, perangkat.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PerangkatDesaRepository) DeletePerangkatDesa(perangkatID int) error {
	stmt, err := pr.DB.Prepare("DELETE FROM perangkat WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(perangkatID)
	if err != nil {
		return err
	}

	return nil
}

func (pr *PerangkatDesaRepository) GetPerangkatDesaByID(perangkatID int) (*models.PerangkatDesa, error) {
	stmt, err := pr.DB.Prepare("SELECT id, nama, jabatan FROM perangkat WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(perangkatID)

	perangkat := &models.PerangkatDesa{}
	err = row.Scan(&perangkat.ID, &perangkat.Nama, &perangkat.Jabatan)
	if err != nil {
		return nil, err
	}

	return perangkat, nil
}

func (pr *PerangkatDesaRepository) GetAllPerangkatDesa() ([]*models.PerangkatDesa, error) {
	stmt, err := pr.DB.Prepare("SELECT id, nama, jabatan FROM perangkat")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	perangkats := []*models.PerangkatDesa{}
	for rows.Next() {
		perangkat := &models.PerangkatDesa{}
		err := rows.Scan(&perangkat.ID, &perangkat.Nama, &perangkat.Jabatan)
		if err != nil {
			return nil, err
		}
		perangkats = append(perangkats, perangkat)
	}

	return perangkats, nil
}