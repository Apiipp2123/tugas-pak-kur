package main

import (
	"coba1/view"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func menuUtama() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah Data Pekerja")
		fmt.Println("2. Tampilkan Data Pekerja")
		fmt.Println("3. Update Data Pekerja")
		fmt.Println("4. Hapus Data Pekerja")
		fmt.Println("Pilih menu: ")

		//Membaca input dari user 
        input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Menghapus karakter newline (\n)

		// Konversi input dari user
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, silahkan masukkan angka.")
			continue
		}

		switch choice{
		case 1:
			fmt.Println("Anda memilih: Tambah Data Pekerja")
			view.Insert()
		case 2:
			fmt.Println("Anda memilih: Tampilkan Data Pekerja")
			view.Views()
		case 3:
			fmt.Println("Anda memilih: Update Data Pekerja")
			view.Update()
		case 4:
			fmt.Println("Anda memilih: Delete Data Pekerja")
			view.Delete()
		case 5:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan Tidak Valid, Silahkan coba lagi")
		}
	}
}

func main() {
	menuUtama()
}