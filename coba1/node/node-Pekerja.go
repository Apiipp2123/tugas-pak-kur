package node

type Address struct {
	Jalan, Kota string
	Nomer int
}

type Pekerja struct {
	ID int
	Nama string
	Alamat Address
	NoTelp string
	Email string
}
