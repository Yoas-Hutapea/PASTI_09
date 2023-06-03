package models

type Kegiatan struct {
	ID           int    `json:"id"`
	Judul        string `json:"judul"`
	Tempat       string `json:"tempat"`
	TanggalMulai string `json:"tanggal_mulai"`
	TanggalAkhir string `json:"tanggal_akhir"`
	Deskripsi    string `json:"deskripsi"`
}
