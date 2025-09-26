# Soal 1: Validasi CLI Sederhana (Input Nama dan Umur)

## Deskripsi
Program ini adalah aplikasi Command Line Interface (CLI) sederhana yang dibuat menggunakan Go. Fungsi utamanya adalah menerima input nama dan umur dari pengguna, kemudian melakukan validasi dasar terhadap umur. Jika umur kurang dari 18 tahun, program akan menampilkan pesan error. Jika umur valid, program akan menampilkan pesan sambutan.

Program ini menggunakan package bawaan Go seperti `bufio` untuk membaca input, `strconv` untuk konversi tipe data, `strings` untuk membersihkan input, dan `fmt` untuk menampilkan output serta membuat error.

## Cara Menjalankan Program

Untuk menjalankan program ini, navigasi ke direktori `soal-1-cli-validasi/` di terminal Anda, lalu gunakan perintah `go run main.go`. Program akan meminta Anda untuk memasukkan nama dan umur secara interaktif.

**Langkah-langkah:**

1.  Buka terminal atau Command Prompt.
2.  Navigasi ke direktori proyek Anda:
    ```bash
    cd path/to/belajar-golang-hsi/soal-1-cli-validasi/
    ```
    *(Ganti `path/to/belajar-golang-hsi` dengan jalur aktual ke folder utama Anda)*

3.  Jalankan program:
    ```bash
    go run main.go
    ```

**Contoh Penggunaan:**

### 1. Umur Valid (Selamat Datang):
```bash
$ go run main.go
Masukan nama Anda:
Budi
Masukan umur Anda:
25
Selamat datang, Budi!