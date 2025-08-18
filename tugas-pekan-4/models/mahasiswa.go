package models

import "gorm.io/gorm"

type Mahasiswa struct {
	ID    uint
	Nama  string
	Tugas []Tugas
}

func SeedMahasiswa(db *gorm.DB) error {
	var count int64
	db.Model(&Mahasiswa{}).Count(&count)
	if count > 0 {
		return nil
	}

	mahasiswaList := []Mahasiswa{
		{Nama: "Andi Pratama"},
		{Nama: "Budi Santoso"},
		{Nama: "Citra Lestari"},
		{Nama: "Dian Kusuma"},
		{Nama: "Eka Sari"},
	}

	return db.Create(&mahasiswaList).Error
}
