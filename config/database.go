package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	usersDBName       = "db_users"
	perangkatDBName   = "db_perangkat"
	kegiatanDBName    = "db_kegiatan"
	pengumumanDBName  = "db_pengumuman"
	mysqlUsername     = "root"
	mysqlPassword     = ""
	mysqlHost         = "localhost"
	mysqlPort         = "3306"
	mysqlMaxOpenConns = 10
	mysqlMaxIdleConns = 5
)

var (
	usersDB       	*sql.DB
	perangkatDB   	*sql.DB
	kegiatanDB    	*sql.DB
	pengumumanDB    *sql.DB
)

// InitializeDB initializes the database connections
func InitializeDB() error {
	usersDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, usersDBName)
	perangkatDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, perangkatDBName)
	kegiatanDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, kegiatanDBName)
	pengumumanDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, pengumumanDBName)

	usersDB, err := sql.Open("mysql", usersDSN)
	if err != nil {
		return err
	}

	perangkatDB, err := sql.Open("mysql", perangkatDSN)
	if err != nil {
		return err
	}

	kegiatanDB, err := sql.Open("mysql", kegiatanDSN)
	if err != nil {
		return err
	}
	
	pengumumanDB, err := sql.Open("mysql", pengumumanDSN)
	if err != nil {
		return err
	}

	usersDB.SetMaxOpenConns(mysqlMaxOpenConns)
	usersDB.SetMaxIdleConns(mysqlMaxIdleConns)

	perangkatDB.SetMaxOpenConns(mysqlMaxOpenConns)
	perangkatDB.SetMaxIdleConns(mysqlMaxIdleConns)

	kegiatanDB.SetMaxOpenConns(mysqlMaxOpenConns)
	kegiatanDB.SetMaxIdleConns(mysqlMaxIdleConns)

	pengumumanDB.SetMaxOpenConns(mysqlMaxOpenConns)
	pengumumanDB.SetMaxIdleConns(mysqlMaxIdleConns)

	err = usersDB.Ping()
	if err != nil {
		return err
	}

	err = perangkatDB.Ping()
	if err != nil {
		return err
	}

	err = kegiatanDB.Ping()
	if err != nil {
		return err
	}

	err = pengumumanDB.Ping()
	if err != nil {
		return err
	}

	return nil
}

// GetUsersDB returns the users database connection
func GetUsersDB() *sql.DB {
	return usersDB
}

// GetPerangkatDB returns the perangkat database connection
func GetPerangkatDB() *sql.DB {
	return perangkatDB
}

// GetKegiatanDB returns the kegiatan database connection
func GetKegiatanDB() *sql.DB {
	return kegiatanDB
}

// GetPengumumanDB returns the pengumuman database connection
func GetPengumumanDB() *sql.DB {
	return pengumumanDB
}