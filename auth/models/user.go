package models

type LoginRequest struct {
	NIK      string `json:"nik"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string     `json:"token"`
	User  UserDetail `json:"user"`
}

type UserDetail struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	NIK           string `json:"nik"`
	NoTelp        string `json:"no_telp"`
	Alamat        string `json:"alamat"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	Usia          int    `json:"usia"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Pekerjaan     string `json:"pekerjaan"`
	Agama         string `json:"agama"`
	KK            string `json:"kk"`
	Gambar        string `json:"gambar"`
}

type User struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	NIK           string `json:"nik"`
	NoTelp        string `json:"no_telp"`
	Alamat        string `json:"alamat"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	Usia          int    `json:"usia"`
	JenisKelamin  string `json:"jenis_kelamin"`
	Pekerjaan     string `json:"pekerjaan"`
	Agama         string `json:"agama"`
	KK            string `json:"kk"`
	Gambar        string `json:"gambar"`
	Password      string `json:"password"`
}
