package repositories

import (
	"database/sql"

	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/models"
)

type KegiatanRepository struct {
	DB *sql.DB
}

func NewKegiatanRepository(dbKegiatan *sql.DB) *KegiatanRepository {
	// initialize and configure the UserRepository instance
	return &KegiatanRepository{
		// initialize fields and dependencies
		DB : dbKegiatan,
	}
}

func (kr *KegiatanRepository) AddKegiatan(kegiatan *models.Kegiatan) error {
	stmt, err := kr.DB.Prepare("INSERT INTO kegiatan (judul, tempat, tanggal_mulai, tanggal_akhir, deskripsi) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kegiatan.Judul, kegiatan.Tempat, kegiatan.TanggalMulai, kegiatan.TanggalAkhir, kegiatan.Deskripsi)
	if err != nil {
		return err
	}

	return nil
}

func (kr *KegiatanRepository) UpdateKegiatan(kegiatan *models.Kegiatan) error {
	stmt, err := kr.DB.Prepare("UPDATE kegiatan SET judul=?, tempat=?, tanggal_mulai=?, tanggal_akhir=?, deskripsi=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kegiatan.Judul, kegiatan.Tempat, kegiatan.TanggalMulai, kegiatan.TanggalAkhir, kegiatan.Deskripsi, kegiatan.ID)
	if err != nil {
		return err
	}

	return nil
}

func (kr *KegiatanRepository) DeleteKegiatan(kegiatanID int) error {
	stmt, err := kr.DB.Prepare("DELETE FROM kegiatan WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kegiatanID)
	if err != nil {
		return err
	}

	return nil
}

func (kr *KegiatanRepository) GetKegiatanByID(kegiatanID int) (*models.Kegiatan, error) {
	stmt, err := kr.DB.Prepare("SELECT id, judul, tempat, tanggal_mulai, tanggal_akhir, deskripsi FROM kegiatan WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(kegiatanID)

	kegiatan := &models.Kegiatan{}
	err = row.Scan(&kegiatan.ID, &kegiatan.Judul, &kegiatan.Tempat, &kegiatan.TanggalMulai, &kegiatan.TanggalAkhir, &kegiatan.Deskripsi)
	if err != nil {
		return nil, err
	}

	return kegiatan, nil
}

func (kr *KegiatanRepository) GetAllKegiatan() ([]*models.Kegiatan, error) {
	stmt, err := kr.DB.Prepare("SELECT id, judul, tempat, tanggal_mulai, tanggal_akhir, deskripsi FROM kegiatan")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kegiatans := []*models.Kegiatan{}
	for rows.Next() {
		kegiatan := &models.Kegiatan{}
		err := rows.Scan(&kegiatan.ID, &kegiatan.Judul, &kegiatan.Tempat, &kegiatan.TanggalMulai, &kegiatan.TanggalAkhir, &kegiatan.Deskripsi)
		if err != nil {
			return nil, err
		}
		kegiatans = append(kegiatans, kegiatan)
	}

	return kegiatans, nil
}