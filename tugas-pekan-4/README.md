# 🧩 Tugas Golang Pertemuan 4 – Worker & Manajemen Nilai Mahasiswa

## 📌 Deskripsi

Program ini mensimulasikan sistem penugasan dan penilaian mahasiswa menggunakan:

- GORM & PostgreSQL
- Goroutine, Channel, WaitGroup
- Worker paralel untuk:
  - Penugasan tugas ke mahasiswa
  - Penilaian nilai secara acak
- Idempotensi dijaga agar tidak duplikasi data

---

---

## 🚀 Cara Menjalankan

### 1. Siapkan Database PostgreSQL
- Bisa lewat Docker, atau install manual
- Pastikan `.env` diisi sesuai database lokal

### 2. Jalankan program

```bash
go run cmd/main.go

