package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {

	mhs1 := mahasiswa.BuatMahasiswa("Budi", 20, 80, 90, 85)
	mhs2 := mahasiswa.BuatMahasiswa("Budi", 20, 80, 90, 85)
	mhs3 := mahasiswa.BuatMahasiswa("Ani", 22, 75, 85, 85)
	getTotalUmur := func(mahasiswaList ...mahasiswa.Deskripsi) int {
		total := 0
		for _, m := range mahasiswaList {
			total += m.GetUmur()
		}
		return total
	}

	fmt.Println(mhs1.Info())
	fmt.Println("Rata-rata nilai:", mhs1.RataRata())
	fmt.Println("Umur:", mhs1.GetUmur())
	fmt.Println("------------------------------")
	fmt.Println(mhs2.Info())
	fmt.Println("Rata-rata nilai:", mhs2.RataRata())
	fmt.Println("Umur:", mhs2.GetUmur())
	fmt.Println("------------------------------")
	fmt.Println(mhs3.Info())
	fmt.Println("Rata-rata nilai:", mhs3.RataRata())
	fmt.Println("Umur:", mhs3.GetUmur())
	fmt.Println("------------------------------")
	fmt.Println("Total Umur Mahasiswa:", getTotalUmur(mhs1, mhs2, mhs3))

	fmt.Println("Nilai maksimum:", mahasiswa.GetMaxNilai())
}
