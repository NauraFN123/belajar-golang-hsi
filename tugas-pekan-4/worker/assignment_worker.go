package worker

import (
	"fmt"
	"math/rand"
	"time"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

func AssignmentWorker(db *gorm.DB, tugasChan chan models.Tugas) {
	rand.Seed(time.Now().UnixNano())
	judulList := []string{
		"Tugas Pemrograman Goroutine",
		"Tugas Implementasi WaitGroup",
		"Tugas Implementasi Mutex",
		"Tugas Implementasi Channel",
	}

	var mahasiswaList []models.Mahasiswa
	db.Find(&mahasiswaList)

	for _, m := range mahasiswaList {
		var existing models.Tugas
		db.Where("mahasiswa_id = ?", m.ID).First(&existing)
		if existing.ID != 0 {
			continue // idempotent
		}
		judul := judulList[rand.Intn(len(judulList))]
		tugas := models.Tugas{
			Judul:       judul,
			Deskripsi:   "Deskripsi tugas " + judul,
			MahasiswaID: m.ID,
		}
		db.Create(&tugas)
		tugasChan <- tugas
		fmt.Printf("Tugas '%s' diberikan ke %s\n", tugas.Judul, m.Nama)
	}
	close(tugasChan)
}
