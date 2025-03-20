package view

import (
	"coba1/model"
	"coba1/node"
	"bufio"
	"fmt"
	"os"
)

func Insert() {
	var kota, nama, notelp, email string
	var id, nomer int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan ID Pekerja: ")
	fmt.Scan(&id)

	fmt.Print("Masukkan Nama Pekerja: ")
	nama, _ = reader.ReadString('\n')
	nama = nama[:len(nama)-1]

	fmt.Print("Masukkan Jalan: ")
	jalan, _ := reader.ReadString('\n')
	jalan = jalan[:len(jalan)-1]

	fmt.Print("Masukkan Nomer Rumah: ")
	fmt.Scan(&nomer)

	fmt.Print("Masukkan Kota: ")
	fmt.Scan(&kota)

	fmt.Print("Masukkan Nomer Telpon: ")
	fmt.Scan(&notelp)

	fmt.Print("Masukkan Email: ")
	fmt.Scan(&email)

	// Create new pekerja
	pekerja := node.Pekerja{
	ID: id,
	Nama: nama,
	Alamat: node.Address{jalan, kota, nomer},
	NoTelp: notelp,
	Email: email,
	}
	
	// insert to DaftarPekerja
	cek := model.CreatePekerja(pekerja)
	if cek {
		fmt.Println("== Pekerja berhasil ditambahkan ==")
	} else {
		fmt.Println("Pekerja gagal ditambahkan")
	}
	fmt.Println()
}

	func Views() {
		fmt.Println("Daftar Pegawai")
		for i, emp := range model.ReadPekerja() {
			fmt.Println("--- Pekerja ke -", i+1, "---")
			fmt.Println("ID Pekerja\t:", emp.ID)
			fmt.Println("Nama Pekerja\t:", emp.Nama)
			fmt.Println("Alamat Pekerja\t\t:", emp.Alamat.Jalan, emp.Alamat.Nomer, emp.Alamat.Kota)
			fmt.Println("Nomer Telepon Pekerja\t:", emp.NoTelp)
			fmt.Println("Email Pekerja\t\t:", emp.Email)
		}
	}
	
	func Update() {
		var id, nomer int
		var nama, kota, notelp, email string
		reader := bufio.NewReader(os.Stdin)
		
		fmt.Println("Masukkan ID Pekerja yang akan di update: ")
		fmt.Scan(&id)

		fmt.Println("Masukkan Nama Pekerja: ")
		nama, _ = reader.ReadString('\n')
		nama = nama[:len(nama)-1]

		fmt.Println("Masukkan Kota: ")
		fmt.Scan(&kota)

		fmt.Println("Masukkan Jalan: ")
		jalan, _ := reader.ReadString('\n')
		jalan = jalan[:len(jalan)-1]

		fmt.Println("Masukkan Nomer Telpon Rumah: ")
		fmt.Scan(&nomer)

		fmt.Println("Masukkan Nomer Telpon: ")
		fmt.Scan(&notelp)

		fmt.Println("Masukkan Email: ")
		fmt.Scan(&email)

		//create new pekerja
		pekerja := node.Pekerja{
			ID: id,
			Nama: nama,
			Alamat: node.Address{jalan, kota, nomer},
			NoTelp: notelp,
			Email: email,
		}
		// update pekerja
		cek := model.UpdatePekerja(pekerja, id)
		if cek {
			fmt.Println("== Pekerja berhasil di update ==")
		} else {
			fmt.Println(" Pekerja gagal diupdate")
		}
		fmt.Println()
	}

	func Delete() {
		var id int
		fmt.Print("Masukkan ID Pekerja yang akan dihapus: ")
		fmt.Scan(&id)

		cek := model.DeletePekerja(id)
		if cek {
			fmt.Println(" == Pekerja berhasil dihapus ==")
		} else {
			fmt.Println("Pekerja gagal dihapus")
		}
		fmt.Println()
	}
