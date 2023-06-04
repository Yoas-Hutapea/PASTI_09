package models

type Pengumuman struct {
	ID         int    `json:"id"`
	Tanggal    string `json:"tanggal"`
	Judul      string `json:"judul"`
	Deskripsi  string `json:"deskripsi"`
}
