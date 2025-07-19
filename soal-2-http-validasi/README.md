# Soal 2: API Handler Sederhana - Validasi Data

## Deskripsi
Program ini adalah server HTTP sederhana yang dibuat menggunakan package `net/http` bawaan Go. Server ini memiliki satu endpoint `/validate` yang berfungsi untuk memvalidasi data yang diterima melalui query parameter.

**Ketentuan Validasi:**
- Menerima query parameter `email` dan `age`.
- `email` tidak boleh kosong.
- `age` harus berupa angka dan berusia minimal 18 tahun.
- Setiap request yang masuk akan dilog.
- Penanganan error menggunakan `fmt.Errorf` untuk wrapping dan `errors.Is` untuk checking.

## Cara Menjalankan Program

1.  Navigasi ke direktori `soal-2-http-validasi/` di terminal Anda:
    ```bash
    cd belajar-golang-hsi/soal-2-http-validasi/
    ```
2.  Jalankan server HTTP:
    ```bash
    go run main.go
    ```
    Server akan berjalan di `http://localhost:8080`.

3.  Gunakan `curl` (atau browser) untuk mengirim request ke endpoint `/validate`.

**Contoh Penggunaan `curl`:**

### 1. Validasi Sukses:
```bash
curl "http://localhost:8080/validate?email=user@example.com&age=25"