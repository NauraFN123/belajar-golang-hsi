package main

// import "fmt"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukan nama Anda: ")
	nama, _ := reader.ReadString('\n')
	fmt.Println("Masukan umur Anda: ")
	umur, _ := reader.ReadString('\n')

	//menghapus karakter whitespace dari input
	nama = strings.TrimSpace(nama)
	umur = strings.TrimSpace(umur)

	//mengkonversi umur dari string ke integer
	umurInt, _ := strconv.Atoi(umur)

	if umurInt < 18 {
		// membuat error dengan pesan umur tidak valid
		err := fmt.Errorf("umur tidak valid (minimal 18 tahun)")

		fmt.Println(err)

	} else {
		fmt.Printf("Selamat datang, %s!", nama)
	}

}
