package repositories

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"

	"github.com/Yoas-Hutapea/Microservice_09/penduduk/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(dbUser *sql.DB) *UserRepository {
	// initialize and configure the UserRepository instance
	return &UserRepository{
		// initialize fields and dependencies
		DB: dbUser,
	}
}

func (ur *UserRepository) AddUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	stmt, err := ur.DB.Prepare("INSERT INTO users (nama, nik, no_telp, alamat, tempat_lahir, tanggal_lahir, usia, jenis_kelamin, pekerjaan, agama, kk, gambar, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nama, user.NIK, user.NoTelp, user.Alamat, user.TempatLahir, user.TanggalLahir, user.Usia, user.JenisKelamin, user.Pekerjaan, user.Agama, user.KK, user.Gambar, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	stmt, err := ur.DB.Prepare("UPDATE users SET nama=?, no_telp=?, alamat=?, tempat_lahir=?, tanggal_lahir=?, usia=?, jenis_kelamin=?, pekerjaan=?, agama=?, kk=?, gambar=?, nik=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nama, user.NoTelp, user.Alamat, user.TempatLahir, user.TanggalLahir, user.Usia, user.JenisKelamin, user.Pekerjaan, user.Agama, user.KK, user.Gambar, user.NIK, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(userID int) error {
	stmt, err := ur.DB.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}