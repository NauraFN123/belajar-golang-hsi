package models

type Student struct {
	ID       int    `json:"id"`
	NIM      string `json:"nim"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	Semester int    `json:"semester"`
}

var Students = []Student{
	{ID: 1, NIM: "2021001", Name: "Budi Santoso", Email: "budi@univ.ac.id", Major: "Teknik Informatika", Semester: 6},
	{ID: 2, NIM: "2021002", Name: "Sari Dewi", Email: "sari@univ.ac.id", Major: "Sistem Informasi", Semester: 4},
}
