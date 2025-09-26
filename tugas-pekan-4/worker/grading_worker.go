package worker

import (
	"fmt"
	"math/rand"
	"time"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

func GradingWorker(db *gorm.DB, tugasChan chan models.Tugas) {
	rand.Seed(time.Now().UnixNano())
	for tugas := range tugasChan {
		var existing models.Hasil
		db.Where("tugas_id = ?", tugas.ID).First(&existing)
		if existing.ID != 0 {
			continue // idempotent
		}
		nilai := rand.Intn(101)
		hasil := models.Hasil{TugasID: tugas.ID, Nilai: nilai}
		db.Create(&hasil)
		var m models.Mahasiswa
		db.First(&m, tugas.MahasiswaID)
		fmt.Printf("Nilai %d diberikan ke %s untuk tugas '%s'\n", nilai, m.Nama, tugas.Judul)
	}
}
