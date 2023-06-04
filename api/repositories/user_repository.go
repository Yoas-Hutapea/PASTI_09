package repositories

import (
	"database/sql"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
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
	stmt, err := ur.DB.Prepare("INSERT INTO users (nama, nik, no_telp, alamat, tempat_lahir, tanggal_lahir, usia, jenis_kelamin, pekerjaan, agama, kk, gambar, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Nama, user.NIK, user.NoTelp, user.Alamat, user.TempatLahir, user.TanggalLahir, user.Usia, user.JenisKelamin, user.Pekerjaan, user.Agama, user.KK, user.Gambar, user.Password)
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

	_, err = stmt.Exec(user.Nama, user.NoTelp, user.Alamat, user.TempatLahir, user.TanggalLahir, user.Usia, user.JenisKelamin, user.Pekerjaan, user.Agama, user.KK, user.Gambar, user.NIK)
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

func (ur *UserRepository) GetUserByNIK(userNIK string) (*models.User, error) {
	stmt, err := ur.DB.Prepare("SELECT id, nama, nik, no_telp, alamat, tempat_lahir, tanggal_lahir, usia, jenis_kelamin, pekerjaan, agama, kk, gambar, password FROM users WHERE nik=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(userNIK)

	user := &models.User{}
	err = row.Scan(&user.ID, &user.Nama, &user.NIK, &user.NoTelp, &user.Alamat, &user.TempatLahir, &user.TanggalLahir, &user.Usia, &user.JenisKelamin, &user.Pekerjaan, &user.Agama, &user.KK, &user.Gambar, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetAllUsers() ([]*models.User, error) {
	stmt, err := ur.DB.Prepare("SELECT id, nama, nik, no_telp, alamat, tempat_lahir, tanggal_lahir, usia, jenis_kelamin, pekerjaan, agama, kk, gambar, password FROM users")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Nama, &user.NIK, &user.NoTelp, &user.Alamat, &user.TempatLahir, &user.TanggalLahir, &user.Usia, &user.JenisKelamin, &user.Pekerjaan, &user.Agama, &user.KK, &user.Gambar, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}