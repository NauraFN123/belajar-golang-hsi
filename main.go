package main

import (
	"fmt"
	"sync"

	"tugas-pertemuan-6-dan-7/config"
	"tugas-pertemuan-6-dan-7/models"
	"tugas-pertemuan-6-dan-7/worker"
)

func main() {
	config.ConnectDB()
	db := config.DB
	db.AutoMigrate(&models.Mahasiswa{}, &models.Tugas{}, &models.Hasil{})

	_ = models.SeedMahasiswa(db)

	tugasChan := make(chan models.Tugas)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		worker.AssignmentWorker(db, tugasChan)
	}()

	go func() {
		defer wg.Done()
		worker.GradingWorker(db, tugasChan)
	}()

	wg.Wait()

	// Tampilkan hasil akhir
	fmt.Println("\nHasil Tugas Mahasiswa:")
	var tugasList []models.Tugas
	db.Preload("Hasil").Find(&tugasList)
	for _, t := range tugasList {
		var m models.Mahasiswa
		db.First(&m, t.MahasiswaID)
		fmt.Printf("%s - %s: %d\n", m.Nama, t.Judul, t.Hasil.Nilai)
	}
}
