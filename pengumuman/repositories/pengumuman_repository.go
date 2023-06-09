package repositories

import (
	"database/sql"

	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/models"
)

type PengumumanRepository struct {
	DB *sql.DB
}

func NewPengumumanRepository(dbPengumuman *sql.DB) *PengumumanRepository {
	// initialize and configure the UserRepository instance
	return &PengumumanRepository{
		// initialize fields and dependencies
		DB: dbPengumuman,
	}
}

func (pg *PengumumanRepository) AddPengumuman(pengumuman *models.Pengumuman) error {
	stmt, err := pg.DB.Prepare("INSERT INTO pengumuman (tanggal, judul, deskripsi) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pengumuman.Tanggal, pengumuman.Judul, pengumuman.Deskripsi)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PengumumanRepository) UpdatePengumuman(pengumuman *models.Pengumuman) error {
	stmt, err := pg.DB.Prepare("UPDATE pengumuman SET tanggal=?, judul=?, deskripsi=? WHERE id=?",)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pengumuman.Tanggal, pengumuman.Judul, pengumuman.Deskripsi, pengumuman.ID)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PengumumanRepository) DeletePengumuman(pengumumanID int) error {
	stmt, err := pg.DB.Prepare("DELETE FROM pengumuman WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pengumumanID)
	if err != nil {
		return err
	}

	return nil
}

func (pg *PengumumanRepository) GetPengumumanByID(pengumumanID int) (*models.Pengumuman, error) {
	stmt, err := pg.DB.Prepare("SELECT id, tanggal, judul, deskripsi FROM pengumuman WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(pengumumanID)

	pengumuman := &models.Pengumuman{}
	err = row.Scan(&pengumuman.ID, &pengumuman.Tanggal, &pengumuman.Judul, &pengumuman.Deskripsi)
	if err != nil {
		return nil, err
	}

	return pengumuman, nil
}

func (pg *PengumumanRepository) GetAllPengumuman() ([]*models.Pengumuman, error) {
	stmt, err := pg.DB.Prepare("SELECT id, tanggal, judul, deskripsi FROM pengumuman")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pengumumans := []*models.Pengumuman{}
	for rows.Next() {
		pengumuman := &models.Pengumuman{}
		err := rows.Scan(&pengumuman.ID, &pengumuman.Tanggal, &pengumuman.Judul, &pengumuman.Deskripsi)
		if err != nil {
			return nil, err
		}
		pengumumans = append(pengumumans, pengumuman)
	}

	return pengumumans, nil
}
