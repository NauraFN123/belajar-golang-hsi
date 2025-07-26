package mahasiswa

func hitungRataRata(nilai ...int) float64 {
	if len(nilai) == 0 {
		return 0
	}
	var total int
	for _, n := range nilai {
		total += n
	}
	return float64(total) / float64(len(nilai))
}

func BuatMahasiswa(nama string, umur int, nilai ...int) *Mahasiswa {
	nilaiAvg := hitungRataRata(nilai...)
	return &Mahasiswa{
		Nama:     nama,
		Nilai:    nilai,
		umur:     umur,
		nilaiAvg: nilaiAvg,
	}
}
