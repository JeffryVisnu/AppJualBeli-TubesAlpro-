package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX int = 100
var modal,sisaModal,hutang,keuntungan int 

type detBarang struct {
	idBarang       int
	hargaJual int
	hargaBeli int
	nama     string
	kategori string
	stock int
}

type detTransaksi struct {
	idTransaksi int
	idBarangTerjual int
	hargaTerjual int
	stockTerjual int
	namaTerjual string
	kategoriTerjual string
}

type barang [NMAX]detBarang
type transaksi [NMAX]detTransaksi

func main() {
	var B barang
	var T transaksi
	var m,n int
	var mainmenu int
	var running bool = true

	for modal <= 0 {
		fmt.Print("Masukan Modal Awal Toko: Rp. ")
		fmt.Scan(&modal)
		if modal <= 0{
			clearScreen()
			fmt.Println("Modal tidak boleh kurang dari atau sama dengan 0!")
		}
	}
	clearScreen()
	sisaModal = modal
	fmt.Println("~~~~~ Selamat Datang di Toko SenyuminAja ~~~~~\n")
	for running {
		fmt.Println("~~~~~ Main Menu ~~~~~")
		fmt.Println("1. Modifikasi Data Barang")
		fmt.Println("2. Modifikasi Data Transaksi")
		fmt.Println("3. Urutkan Data")
		fmt.Println("4. Pencarian Data")
		fmt.Println("5. Tampil Data")
		fmt.Println("6. Keluar")
		fmt.Print("Masukan Pilihan (1/2/3/4/5/6): ")
		fmt.Scan(&mainmenu)
		clearScreen()
		if mainmenu == 1 {
			modifDataBarang(&B, &n)
		} else if mainmenu == 2 {
			modifDataTransaksi(&B, &T, &n, &m)
		} else if mainmenu == 3 {
			menuUrutkanData(B, T, n, m)
		} else if mainmenu == 4 {
			menuCari(B, T, n, m)
		} else if mainmenu == 5 {
			menuTampilData(B, T, n, m)
		} else if mainmenu == 6 {
			clearScreen()
			fmt.Println("Terimakasih!")
			running = false
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
		}
	}
}

// FUNGSI-FUNGSI MEMODIFIKASI DATA BARANG
func modifDataBarang(B *barang, n *int) {
	var pilih int
	fmt.Println("~~~~~ Menu Modifikasi Data Barang ~~~~~")
	fmt.Println("1. Lihat Data Barang")
	fmt.Println("2. Menambahkan Barang")
	fmt.Println("3. Mengedit Barang")
	fmt.Println("4. Menghapus Barang")
	fmt.Println("5. Kembali")
	fmt.Print("Masukan Pilihan (1/2/3/4/5): ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		clearScreen()
		tampilBarang(*B,*n)
		modifDataBarang(B,n)
	} else if pilih == 2 {
		clearScreen()
		menambahBarang(B, n)
		modifDataBarang(B,n)
	} else if pilih == 3 {
		clearScreen()
		mengeditBarang(B, *n)
		modifDataBarang(B,n)
	} else if pilih == 4 {
		clearScreen()
		menghapusBarang(B, n)
		modifDataBarang(B,n)
	} else if pilih == 5 {
		clearScreen()
		return
	} else {
		clearScreen()
		fmt.Println("Masukan tidak valid!\n")
		modifDataBarang(B, n)
	}
}

func menambahBarang(B *barang, n *int) {
	if *n >= NMAX {
		fmt.Println("~~~ Tambah Barang ~~~")
		fmt.Println("Barang penuh, tidak bisa menambah barang lagi!")
		return
	}
	fmt.Println("~~~ Tambah Barang ~~~")
	fmt.Println("Sisa Modal : ", sisaModal)
	fmt.Println("\n~ Masukan data barang baru ~")
	B[*n].idBarang = *n+1
	validName := false
    for !validName {
        fmt.Print("Nama Barang: ")
        fmt.Scan(&B[*n].nama)
        if CekNamaDiArray(B, *n, B[*n].nama) {
            fmt.Println("\nNama barang sudah ada!")
            fmt.Println("~ Masukkan nama barang yang berbeda ~")
        } else {
            validName = true
        }
	}
	fmt.Print("Kategori Barang: ")
	fmt.Scan(&B[*n].kategori)
	for B[*n].nama == B[*n].kategori {
		fmt.Println("\nNama dan kategori barang tidak boleh sama!")
		fmt.Println("~ Masukan ulang nama dan kategori barang ~")
		fmt.Print("Nama Barang: ")
		fmt.Scan(&B[*n].nama)
		fmt.Print("Kategori Barang: ")
		fmt.Scan(&B[*n].kategori)
	}
	fmt.Print("Harga Beli Barang: Rp. ")
	fmt.Scan(&B[*n].hargaBeli)
	fmt.Print("Harga Jual Barang: Rp. ")
	fmt.Scan(&B[*n].hargaJual)
	for B[*n].hargaBeli >= B[*n].hargaJual {
		fmt.Println("\nHarga jual barang harus lebih tinggi dari harga beli!")
		fmt.Println("~ Masukan ulang harga ~")
		fmt.Print("Harga Beli Barang: Rp. ")
		fmt.Scan(&B[*n].hargaBeli)
		fmt.Print("Harga Jual Barang: Rp. ")
		fmt.Scan(&B[*n].hargaJual)
	}
	fmt.Print("Stock Barang: ")
	fmt.Scan(&B[*n].stock)
	for B[*n].stock <=0 {
		fmt.Println("\nStock barang harus lebih besar dari 0!")
		fmt.Println("~ Masukan ulang stock ~")
		fmt.Print("Stock Barang: ")
		fmt.Scan(&B[*n].stock)
	}
	clearScreen()
	sisaModal = sisaModal - (B[*n].hargaBeli*B[*n].stock)
	if sisaModal < 0 {
		hutang = (sisaModal *-1) + hutang
		sisaModal = 0
		fmt.Println("Modal telah habis, hutang telah ditambahkan!")
	}
	*n++
	fmt.Println("Barang Berhasil ditambahkan!")
	tampilBarang(*B, *n)
}

func CekNamaDiArray(B *barang, n int, name string) bool {
    for i := 0; i < n; i++ {
        if B[i].nama == name {
            return true
        }
    }
    return false
}

func mengeditBarang(B *barang, n int) {
	var editID,menuEdit int
	var editNama,namaBaru string
	var cek int
	var beliBaru,stockBaru int
	fmt.Println("~~~~ Edit Barang ~~~~")
	if n == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilBarang(*B, n)
		fmt.Println("~~ Menu Edit Barang ~~")
		fmt.Println("1. Edit barang menggunakan ID")
		fmt.Println("2. Edit barang menggunakan Nama Barang")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan Pilihan (1/2/3): ")
		fmt.Scan(&menuEdit)
		if menuEdit == 1{ 
			fmt.Print("Masukan ID barang yang akan diedit: ")
			fmt.Scan(&editID)
			cek = cariBarangId(*B,n,editID)
			if cek != -1 {
				clearScreen()
				fmt.Println("Barang ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
				fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				fmt.Println("\n~ Masukan data barang baru ~")
				validName := false
				for !validName {
					fmt.Print("Nama Barang: ")
					fmt.Scan(&namaBaru)
					if CekNamaDiArray(B, n, namaBaru) {
						fmt.Println("\nNama barang sudah ada!")
						fmt.Println("~ Masukkan nama barang yang berbeda ~")
					} else {
						validName = true
						B[cek].nama = namaBaru
					}
				}
				fmt.Print("Kategori Barang: ")
				fmt.Scan(&B[cek].kategori)
				for B[cek].nama == B[cek].kategori {
					fmt.Println("\nNama dan kategori barang tidak boleh sama!")
					fmt.Println("~ Masukan ulang nama dan kategori barang ~")
					fmt.Print("Nama Barang: ")
					fmt.Scan(&B[cek].nama)
					fmt.Print("Kategori Barang: ")
					fmt.Scan(&B[cek].kategori)
				}
				fmt.Print("Harga Beli Barang: Rp. ")
				fmt.Scan(&beliBaru)
				fmt.Print("Harga Jual Barang: Rp. ")
				fmt.Scan(&B[cek].hargaJual)
				for beliBaru >= B[cek].hargaJual {
					fmt.Println("\nHarga jual barang harus lebih tinggi dari harga beli!")
					fmt.Println("~ Masukan ulang harga ~")
					fmt.Print("Harga Beli Barang: Rp. ")
					fmt.Scan(&beliBaru)
					fmt.Print("Harga Jual Barang: Rp. ")
					fmt.Scan(&B[cek].hargaJual)
				}
				fmt.Print("Persediaan Barang: ")
				fmt.Scan(&stockBaru)
				for stockBaru <=0 {
					fmt.Println("\nStock barang harus lebih besar dari 0!")
					fmt.Println("~ Masukan ulang stock ~")
					fmt.Print("Stock Barang: ")
					fmt.Scan(&stockBaru)
				}
				if beliBaru*stockBaru > B[cek].hargaBeli*B[cek].stock {
					sisaModal = sisaModal - (beliBaru*stockBaru - B[cek].hargaBeli*B[cek].stock)
				} else if beliBaru*stockBaru < B[cek].hargaBeli*B[cek].stock {
					sisaModal = sisaModal + (B[cek].hargaBeli*B[cek].stock - beliBaru*stockBaru)
				}
				B[cek].hargaBeli = beliBaru
				B[cek].stock = stockBaru
				if sisaModal < 0 {
					hutang = (sisaModal *-1) + hutang
					sisaModal = 0
				}
				clearScreen()
				fmt.Println("Data barang berhasil diedit!")
				tampilBarang(*B,n)
			} else {
				clearScreen()
				fmt.Println("Barang tidak dapat ditemukan!\n")
				mengeditBarang(B,n)
			}
		} else if menuEdit== 2 {
			fmt.Print("Masukan nama barang yang akan diedit: ")
			fmt.Scan(&editNama)
			cek = cariBarangNama(*B,n,editNama)
			if cek != -1 {
				clearScreen()
				fmt.Println("Barang ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
				fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				fmt.Println("\n~ Masukan data barang baru ~")
				validName := false
				for !validName {
					fmt.Print("Nama Barang: ")
					fmt.Scan(&namaBaru)
					if CekNamaDiArray(B, n, namaBaru) {
						fmt.Println("\nNama barang sudah ada!")
						fmt.Println("~ Masukkan nama barang yang berbeda ~")
					} else {
						validName = true
						B[cek].nama = namaBaru
					}
				}
				fmt.Print("Kategori Barang: ")
				fmt.Scan(&B[cek].kategori)
				for B[cek].nama == B[cek].kategori {
					fmt.Println("\nNama dan kategori barang tidak boleh sama!")
					fmt.Println("~ Masukan ulang nama dan kategori barang ~")
					fmt.Print("Nama Barang: ")
					fmt.Scan(&B[cek].nama)
					fmt.Print("Kategori Barang: ")
					fmt.Scan(&B[cek].kategori)
				}
				fmt.Print("Harga Beli Barang: Rp. ")
				fmt.Scan(&beliBaru)
				fmt.Print("Harga Jual Barang: Rp. ")
				fmt.Scan(&B[cek].hargaJual)
				for beliBaru >= B[cek].hargaJual {
					fmt.Println("\nHarga jual barang harus lebih tinggi dari harga beli!")
					fmt.Println("~ Masukan ulang harga ~")
					fmt.Print("Harga Beli Barang: Rp. ")
					fmt.Scan(&beliBaru)
					fmt.Print("Harga Jual Barang: Rp. ")
					fmt.Scan(&B[cek].hargaJual)
				}
				fmt.Print("Persediaan Barang: ")
				fmt.Scan(&stockBaru)
				for stockBaru <=0 {
					fmt.Println("\nStock barang harus lebih besar dari 0!")
					fmt.Println("~ Masukan ulang stock ~")
					fmt.Print("Stock Barang: ")
					fmt.Scan(&stockBaru)
				}
				if beliBaru*stockBaru > B[cek].hargaBeli*B[cek].stock {
					sisaModal = sisaModal - (beliBaru*stockBaru - B[cek].hargaBeli*B[cek].stock)
				} else if beliBaru*stockBaru < B[cek].hargaBeli*B[cek].stock {
					sisaModal = sisaModal + (B[cek].hargaBeli*B[cek].stock - beliBaru*stockBaru)
				}
				B[cek].hargaBeli = beliBaru
				B[cek].stock = stockBaru
				if sisaModal < 0 {
					hutang = (sisaModal *-1) + hutang
					sisaModal = 0
				}
				clearScreen()
				fmt.Println("Data barang berhasil diedit!")
				tampilBarang(*B,n)
			} else {
				clearScreen()
				fmt.Println("Barang tidak dapat ditemukan!\n")
				mengeditBarang(B,n)
			}
		} else if menuEdit == 3{
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			mengeditBarang(B, n)
		}
	}
}

func menghapusBarang(B *barang, n *int) {
	var hapusID,menuHapus int
	var hapusNama,yakinHapus string
	var cek int
	fmt.Println("~~~~ Hapus Barang ~~~~")
	if *n == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilBarang(*B, *n)
		fmt.Println("~~ Menu Hapus Barang ~~")
		fmt.Println("1. Hapus barang menggunakan ID")
		fmt.Println("2. Hapus barang menggunakan Nama Barang")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan Pilihan (1/2/3): ")
		fmt.Scan(&menuHapus)
		if menuHapus == 1{ 
			fmt.Print("Masukan ID barang yang akan dihapus: ")
			fmt.Scan(&hapusID)
			cek = cariBarangId(*B,*n,hapusID)
			if cek != -1 {
				clearScreen()
				fmt.Println("Barang ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
				fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				fmt.Print("Apakah anda yakin akan menghapus data barang tersebut?(Y/N): ")
				fmt.Scan(&yakinHapus)
				if yakinHapus == "Y" || yakinHapus == "y" {
				sisaModal = sisaModal+ (B[cek].hargaBeli*B[cek].stock)+hutang
				for i:=cek;i<*n-1;i++{
					B[i]=B[i+1]
					B[i].idBarang = i + 1
				}
				*n--
				clearScreen()
				fmt.Println("Data barang berhasil dihapus!\n")
				tampilBarang(*B,*n)
				} else if yakinHapus == "N" || yakinHapus == "n"{
					clearScreen()
					fmt.Println("Barang tidak dihapus!")
					tampilBarang(*B,*n)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					menghapusBarang(B,n)
				}
			} else {
				clearScreen()
				fmt.Println("Barang tidak dapat ditemukan!\n")
				menghapusBarang(B,n)
			}
		} else if menuHapus== 2 {
			fmt.Print("Masukan nama barang yang akan dihapus: ")
			fmt.Scan(&hapusNama)
			cek = cariBarangNama(*B,*n,hapusNama)
			if cek != -1 {
				clearScreen()
				fmt.Println("Barang ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
				fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				fmt.Print("Apakah anda yakin akan menghapus data barang tersebut?(Y/N): ")
				fmt.Scan(&yakinHapus)
				if yakinHapus == "Y" || yakinHapus == "y" {
				sisaModal = sisaModal+ (B[cek].hargaBeli*B[cek].stock)+hutang
				for i:=cek;i<*n-1;i++{
					B[i]=B[i+1]
					B[i].idBarang = i + 1
				}
				*n--
				clearScreen()
				fmt.Println("Data barang berhasil dihapus!\n")
				tampilBarang(*B,*n)
				} else if yakinHapus == "N" || yakinHapus == "n"{
					clearScreen()
					fmt.Println("Barang tidak dihapus!")
					tampilBarang(*B,*n)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					menghapusBarang(B,n)
				}
			} else {
				clearScreen()
				fmt.Println("\nBarang tidak dapat ditemukan!\n")
			}
		} else if menuHapus == 3 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menghapusBarang(B, n)
		}
	}
}

// FUNGSI-FUNGSI MEMODIFIKASI DATA TRANSAKSI
func modifDataTransaksi(B *barang, T *transaksi, n, m *int) {
		var pilih int
		fmt.Println("~~~~~ Menu Tampil Data Transaksi ~~~~~")
		if *n <= 0 {
				fmt.Println("Tidak ada barang yang tersedia!\n")
			} else {
			fmt.Println("1. Lihat Data Transaksi")
			fmt.Println("2. Menambahkan Data Transaksi")
			fmt.Println("3. Mengedit Data Transaksi")
			fmt.Println("4. Menghapus Data Transaksi")	
			fmt.Println("5. Kembali")
			fmt.Print("Masukan Pilihan (1/2/3/4/5): ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				clearScreen()
				tampilDataTransaksi(*T,*m)
				modifDataTransaksi(B,T,n,m)
			} else if pilih == 2 {
				clearScreen()
				menuMenambahDataTransaksi(B, T, n,m)
				modifDataTransaksi(B,T,n,m)
			} else if pilih == 3 {
				clearScreen()
				menuMengeditDataTransaksi(B,T,n,m)
				modifDataTransaksi(B,T,n,m)
			} else if pilih == 4 {
				clearScreen()
				menuMenghapusDataTransaksi(B,T,n,m)
				modifDataTransaksi(B,T,n,m)
			} else if pilih == 5 {
				clearScreen()
				return
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				modifDataTransaksi(B,T,n,m)
			}
		}
}

func menuMenambahDataTransaksi(B *barang, T *transaksi, n,m *int){
	var pilih int
	fmt.Println("~~~ Menu Tambah Data Transaksi ~~~")
	tampilBarang(*B,*n)
	fmt.Println("1. Tambah data transaksi dengan ID barang")
	fmt.Println("2. Tambah data transaksi dengan nama barang")
	fmt.Println("3. Kembali")
	fmt.Print("Masukan Pilihan (1/2/3): ")	
	fmt.Scan(&pilih)
	if pilih == 1 {
		clearScreen()
		tambahDataTransaksiID(B,T,n,m)
	} else if pilih == 2 {
		clearScreen()
		tambahDataTransaksiNama(B,T,n,m)
	} else if pilih == 3 {
		clearScreen()
		return
	} else {
		clearScreen()
		fmt.Println("Masukan tidak valid!\n")
		menuMenambahDataTransaksi(B,T,n,m)
	}
}

func tambahDataTransaksiID(B *barang, T *transaksi, n ,m*int){
	var idJual,banyakterjual,cek int
	fmt.Println("~~~ Tambah Data Transaksi dengan ID Barang ~~~")
	fmt.Println("\n~ Masukan data transaksi baru ~")
	T[*m].idTransaksi = *m+1
	fmt.Print("Masukan ID barang yang akan dijual: ")
	fmt.Scan(&idJual)
	cek = cariBarangId(*B,*n,idJual)
		if cek != -1 {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock")
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaJual, B[cek].stock)
			fmt.Print("Masukan banyak barang yang akan dijual: ")
			fmt.Scan(&banyakterjual)
			if banyakterjual<=B[cek].stock {
				T[*m].idBarangTerjual = cek+1
				T[*m].namaTerjual = B[cek].nama
				T[*m].kategoriTerjual = B[cek].kategori
				T[*m].hargaTerjual = B[cek].hargaJual
				T[*m].stockTerjual = banyakterjual
				B[cek].stock = B[cek].stock-banyakterjual
				keuntungan = (B[cek].hargaJual * banyakterjual)+keuntungan
				clearScreen()
				fmt.Println("Penjualan barang berhasil!\n")
				*m++
				tampilDataTransaksi(*T,*m)
			} else {
				clearScreen()
				fmt.Println("Banyak barang yang akan dijual melebihi persediaan yang ada!\n")
				menuMenambahDataTransaksi(B,T,n,m)
			}
		} else {
			clearScreen()
			fmt.Println("ID barang tidak dapat di temukan!\n")
			menuMenambahDataTransaksi(B,T,n,m)
		}
}

func tambahDataTransaksiNama(B *barang,T *transaksi, n,m *int){
	var banyakterjual,cek int
	var namaJual string
	fmt.Println("~~~ Tambah Data Transaksi dengan Nama Barang ~~~")
	fmt.Println("\n~ Masukan data transaksi baru ~")
	T[*m].idTransaksi = *m+1
	fmt.Println("Masukan Nama barang yang akan dijual: ")
	fmt.Scan(&namaJual)
	cek = cariBarangNama(*B,*n,namaJual)
		if cek != -1 {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock")
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaJual, B[cek].stock)
			fmt.Print("Masukan banyak barang yang akan dijual: ")
			fmt.Scan(&banyakterjual)
			if banyakterjual<=B[cek].stock {
				T[*m].idBarangTerjual = cek+1
				T[*m].namaTerjual = B[cek].nama
				T[*m].kategoriTerjual = B[cek].kategori
				T[*m].hargaTerjual = B[cek].hargaJual
				T[*m].stockTerjual = banyakterjual
				B[cek].stock = B[cek].stock-banyakterjual
				keuntungan = (B[cek].hargaJual * banyakterjual)+keuntungan
				clearScreen()
				fmt.Println("Penjualan barang berhasil!\n")
				*m++
				tampilDataTransaksi(*T,*m)
			} else {
				clearScreen()
				fmt.Println("Banyak barang yang akan dijual melebihi persediaan yang ada!\n")
				tambahDataTransaksiID(B,T,n,m)
			}
		} else {
			clearScreen()
			fmt.Println("Nama barang tidak dapat di temukan!\n")
			menuMenambahDataTransaksi(B,T,n,m)
		}
}

func menuMengeditDataTransaksi(B *barang,T *transaksi, n,m *int){
	var pilih int
	fmt.Println("~~~ Menu Edit Data Transaksi ~~~")
	if *m == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilDataTransaksi(*T,*m)
		fmt.Println("~~ Edit Transaksi ~~")
		fmt.Println("1. Edit data transaksi dengan ID transaksi")
		fmt.Println("2. Kembali")
		fmt.Print("Masukan Pilihan (1/2): ")	
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			editDataTransaksi(B,T,n,m)
		} else if pilih == 2 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuMengeditDataTransaksi(B,T,n,m)
		}
	}
}

func editDataTransaksi(B *barang,T *transaksi, n,m *int){
	var editID,menuEdit,stockBaru,x int
	var cek int
	fmt.Println("~~~~ Edit Data Transaksi ~~~~")
	if *m == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilDataTransaksi(*T,*m)
		fmt.Println("~~ Menu Edit Data Transaksi ~~")
		fmt.Println("1. Edit transaksi (Stock Terjual) menggunakan ID")
		fmt.Println("2. Kembali")
		fmt.Print("Masukan Pilihan (1/2): ")
		fmt.Scan(&menuEdit)
		if menuEdit == 1{ 
			fmt.Print("Masukan ID transaksi yang akan diedit: ")
			fmt.Scan(&editID)
			cek = cariTransaksiId(*T,*m,editID)
			if cek != -1 {
				x = cariBarangId(*B,*n,T[cek].idBarangTerjual)
				fmt.Print("Masukan stok barang yang akan dijual: ")
				fmt.Scan(&stockBaru)
				if T[cek].stockTerjual > stockBaru && stockBaru<=B[x].stock+T[cek].stockTerjual{ 
					B[x].stock = B[x].stock + (T[cek].stockTerjual-stockBaru) 
					keuntungan = keuntungan-(T[cek].stockTerjual - stockBaru)*T[cek].hargaTerjual
					T[cek].stockTerjual = stockBaru 
					clearScreen()
					fmt.Println("Data Transaksi berhasil diedit!")
				} else if T[cek].stockTerjual < stockBaru && stockBaru<=B[x].stock+T[cek].stockTerjual {
					B[x].stock = B[x].stock - (stockBaru-T[cek].stockTerjual)
					keuntungan = keuntungan+(stockBaru-T[cek].stockTerjual)*T[cek].hargaTerjual
					T[cek].stockTerjual = stockBaru
					clearScreen()
					fmt.Println("Data Transaksi berhasil diedit!\n")
					tampilDataTransaksi(*T,*m)
				} else {
					clearScreen()
					fmt.Println("Stock yang tersedia hanya tersisa",B[x].stock)
					editDataTransaksi(B,T,n,m)
				}
			} else {
				clearScreen()
				fmt.Println("Data transaksi tidak dapat ditemukan!")
				editDataTransaksi(B,T,n,m)
			}
		} else if menuEdit == 2{
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			editDataTransaksi(B,T,n,m)
		}
	}
}

func menuMenghapusDataTransaksi(B *barang,T *transaksi, n,m *int){
	var pilih int
	fmt.Println("~~~ Menu Hapus Data Transaksi ~~~")
	if *m == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilDataTransaksi(*T,*m)
		fmt.Println("1. Hapus data transaksi dengan ID transaksi")
		fmt.Println("2. Kembali")
		fmt.Print("Masukan Pilihan (1/2): ")	
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			hapusDataTransaksi(B,T,n,m)
		} else if pilih == 2 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuMenghapusDataTransaksi(B,T,n,m)
		}
	}
}

func hapusDataTransaksi(B *barang,T *transaksi, n,m *int){
	var hapusID,menuHapus,x int
	var yakinHapus string
	var cek int
	fmt.Println("~~~~ Hapus Data Transaksi ~~~~")
	if *m == 0 {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else { 
		tampilDataTransaksi(*T,*m)
		fmt.Println("~~ Menu Hapus Data Transaksi ~~")
		fmt.Println("1. Hapus transaksi menggunakan ID")
		fmt.Println("2. Kembali")
		fmt.Print("Masukan Pilihan (1/2): ")
		fmt.Scan(&menuHapus)
		if menuHapus == 1{ 
			fmt.Print("Masukan ID transaksi yang akan dihapus: ")
			fmt.Scan(&hapusID)
			cek = cariTransaksiId(*T,*m,hapusID )
			if cek != -1 {
				x = cariBarangId(*B,*n,T[cek].idBarangTerjual)
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock Terjual")
				fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[cek].idTransaksi, T[cek].idBarangTerjual, T[cek].namaTerjual, T[cek].kategoriTerjual, T[cek].hargaTerjual, T[cek].stockTerjual)
				fmt.Print("Apakah anda yakin akan menghapus data transaksi tersebut?(Y/N) :")
				fmt.Scan(&yakinHapus)
				if yakinHapus == "Y" || yakinHapus == "y"{
					keuntungan= keuntungan - (T[cek].hargaTerjual*T[cek].stockTerjual)
					B[x].stock = B[x].stock + T[cek].idBarangTerjual
					if cek != -1 {
						for i:=cek;i<*m-1;i++{
							T[i]=T[i+1]
							T[i].idTransaksi = i + 1
						}
					}
					*m--
					clearScreen()
					fmt.Println("Transaksi berhasil dihapus!\n")
					tampilDataTransaksi(*T,*m)
				} else if yakinHapus == "N" || yakinHapus == "n" {
					clearScreen()
					fmt.Println("Transaksi tidak dihapus!")
					menuMenghapusDataTransaksi(B,T,n,m)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					hapusDataTransaksi(B,T,n,m)
				}
			} else {
				clearScreen()
				fmt.Println("Data transaksi tidak dapat ditemukan!")
				hapusDataTransaksi(B,T,n,m)
			}
		} else if menuHapus == 2{
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			hapusDataTransaksi(B,T,n,m)
		}
	}
}


// FUNGSI-FUNGSI MENGURUTKAN DATA
func menuUrutkanData(B barang,T transaksi, n,m int) {
	var pilih int
	fmt.Println("~~~~~ Menu Mengurutkan Data ~~~~~")
		fmt.Println("1. Urutkan Data Barang")
		fmt.Println("2. Urutkan Data Transaksi")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan pilihan (1/2/3): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			menuUrutkanDataBarang(&B,&T,&n,&m)
		} else if pilih == 2 {
			clearScreen()
			menuUrutkanDataTransaksi(&T,&m)
		} else if pilih == 3 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuUrutkanData(B,T,n,m)
	}
}

// URUTKAN DATA BARANG
func menuUrutkanDataBarang(B *barang,T *transaksi, n,m *int){
	var pilih,urutkan int
	fmt.Println("~~~~ Urutkan Data Barang ~~~~")
	if *n <= 0  {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else {
		fmt.Println("~~ Menu Urutkan Data Barang ~~")
		fmt.Println("1. Urutkan Data Barang berdasarkan ID Barang")
		fmt.Println("2. Urutkan Data Barang berdasarkan Harga Beli Barang")
		fmt.Println("3. Urutkan Data Barang berdasarkan Harga Jual Barang")
		fmt.Println("4. Urutkan Data Barang berdasarkan Stock Barang")
		fmt.Println("5. Kembali")
		fmt.Print("Masukan pilihan (1/2/3/4/5): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			fmt.Println("Urutkan Data Barang berdasarkan ID Barang secara?")
			fmt.Println("1. Membesar (Ascending)")
			fmt.Println("2. Mengecil (Descending)")
			fmt.Print("Masukan pilihan (1/2): ")
			fmt.Scan(&urutkan)
			if urutkan == 1 {
				clearScreen()
				urutkanIdBarangAsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else if urutkan == 2{
				clearScreen()
				urutkanIdBarangDsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				menuUrutkanDataBarang(B,T,n,m)
			}
		} else if pilih == 2 {
			clearScreen()
			fmt.Println("Urutkan Data Barang berdasarkan Harga Beli Barang secara?")
			fmt.Println("1. Membesar (Ascending)")
			fmt.Println("2. Mengecil (Descending)")
			fmt.Print("Masukan pilihan (1/2): ")
			fmt.Scan(&urutkan)
			if urutkan == 1 {
				clearScreen()
				urutkanHargaBeliBarangAsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else if urutkan == 2{
				clearScreen()
				urutkanHargaBeliBarangDsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				menuUrutkanDataBarang(B,T,n,m)
			}
		} else if  pilih == 3{
			clearScreen()
			fmt.Println("Urutkan Data Barang berdasarkan Harga Jual Barang secara?")
			fmt.Println("1. Membesar (Ascending)")
			fmt.Println("2. Mengecil (Descending)")
			fmt.Print("Masukan pilihan (1/2): ")
			fmt.Scan(&urutkan)
			if urutkan == 1 {
				clearScreen()
				urutkanHargaJualBarangAsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else if urutkan == 2{
				clearScreen()
				urutkanHargaJualBarangDsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				menuUrutkanDataBarang(B,T,n,m)
			}
		} else if pilih == 4 {
			clearScreen()
			fmt.Println("Urutkan Data Barang berdasarkan Stock Barang secara?")
			fmt.Println("1. Membesar (Ascending)")
			fmt.Println("2. Mengecil (Descending)")
			fmt.Print("Masukan pilihan (1/2): ")
			fmt.Scan(&urutkan)
			if urutkan == 1 {
				clearScreen()
				urutkanStockBarangAsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else if urutkan == 2{
				clearScreen()
				urutkanStockBarangDsc(B,n)
				tampilBarang(*B,*n)
				menuUrutkanDataBarang(B,T,n,m)
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				menuUrutkanDataBarang(B,T,n,m)
				menuUrutkanDataBarang(B,T,n,m)
			}
		}else if pilih == 5 {
			clearScreen()
			menuUrutkanData(*B,*T,*n,*m)
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuUrutkanDataBarang(B,T,n,m)
		}
	}
}

func urutkanIdBarangAsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].idBarang > key.idBarang {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanIdBarangDsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].idBarang < key.idBarang {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanHargaBeliBarangAsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].hargaBeli > key.hargaBeli {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanHargaBeliBarangDsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].hargaBeli < key.hargaBeli {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanHargaJualBarangAsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].hargaJual > key.hargaJual {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanHargaJualBarangDsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].hargaJual < key.hargaJual {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanStockBarangAsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].stock > key.stock {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}

func urutkanStockBarangDsc(B *barang, n *int){
	for i := 1; i < *n; i++ {
        key := B[i]
        j := i - 1
        for j >= 0 && B[j].stock < key.stock {
            B[j+1] = B[j]
            j = j - 1
        }
        B[j+1] = key
    }
}


// URUTKAN DATA TRANSAKSI
func menuUrutkanDataTransaksi(T *transaksi, m *int){
	var pilih,urutkan int
	if *m <= 0  {
		fmt.Println("Tidak ada transaksi yang tersedia!\n")
	} else {
		fmt.Println("~~~~~ Urutkan Data Transaksi ~~~~~")
			fmt.Println("~~ Menu Urutkan Data Transaksi ~~")
			fmt.Println("1. Urutkan Data Transaksi berdasarkan ID Transaksi")
			fmt.Println("2. Urutkan Data Transaksi berdasarkan Harga Jual Barang")
			fmt.Println("3. Urutkan Data Transaksi berdasarkan Stock Terjual")
			fmt.Println("4. Kembali")
			fmt.Print("Masukan pilihan (1/2/3/4): ")
			fmt.Scan(&pilih)
			if pilih == 1 {
				clearScreen()
				fmt.Println("Urutkan Data Transaksi berdasarkan ID Transaksi secara?")
				fmt.Println("1. Membesar (Ascending)")
				fmt.Println("2. Mengecil (Descending)")
				fmt.Print("Masukan pilihan (1/2): ")
				fmt.Scan(&urutkan)
				if urutkan == 1 {
					clearScreen()
					urutkanIdTransAsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else if urutkan == 2{
					clearScreen()
					urutkanIdTransDsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					menuUrutkanDataTransaksi(T,m)
				}
			} else if pilih == 2 {
				clearScreen()
				fmt.Println("Urutkan Data Transaksi berdasarkan Harga Jual Barang secara?")
				fmt.Println("1. Membesar (Ascending)")
				fmt.Println("2. Mengecil (Descending)")
				fmt.Print("Masukan pilihan (1/2): ")
				fmt.Scan(&urutkan)
				if urutkan == 1 {
					clearScreen()
					urutkanHargaJualTransaksiAsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else if urutkan == 2{
					clearScreen()
					urutkanHargaJualTransaksiDsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					menuUrutkanDataTransaksi(T,m)
				}
			} else if  pilih == 3{
				clearScreen()
				fmt.Println("Urutkan Data Transaksi berdasarkan Stock Terjual secara?")
				fmt.Println("1. Membesar (Ascending)")
				fmt.Println("2. Mengecil (Descending)")
				fmt.Print("Masukan pilihan (1/2): ")
				fmt.Scan(&urutkan)
				if urutkan == 1 {
					clearScreen()
					urutkanStockTerjualAsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else if urutkan == 2{
					clearScreen()
					urutkanStockTerjualDsc(T,m)
					tampilDataTransaksi(*T,*m)
				} else {
					clearScreen()
					fmt.Println("Masukan tidak valid!\n")
					menuUrutkanDataTransaksi(T,m)
				}
			} else if pilih == 4 {
				clearScreen()
				return
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid!\n")
				menuUrutkanDataTransaksi(T,m)
		}
	}
}

func urutkanIdTransAsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        minIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].idTransaksi < T[minIndex].idTransaksi {
                minIndex = j
            }
        }
        T[i], T[minIndex] = T[minIndex], T[i]
    }
}

func urutkanIdTransDsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        maxIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].idTransaksi > T[maxIndex].idTransaksi {
                maxIndex = j
            }
        }
        T[i], T[maxIndex] = T[maxIndex], T[i]
    }
}

func urutkanStockTerjualAsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        minIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].stockTerjual < T[minIndex].stockTerjual {
                minIndex = j
            }
        }
        T[i], T[minIndex] = T[minIndex], T[i]
    }
}

func urutkanStockTerjualDsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        maxIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].stockTerjual > T[maxIndex].stockTerjual {
                maxIndex = j
            }
        }
        T[i], T[maxIndex] = T[maxIndex], T[i]
    }
}

func urutkanHargaJualTransaksiAsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        minIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].hargaTerjual < T[minIndex].hargaTerjual {
                minIndex = j
            }
        }
        T[i], T[minIndex] = T[minIndex], T[i]
    }
}

func urutkanHargaJualTransaksiDsc(T *transaksi, m *int){
	for i := 0; i < *m-1; i++ {
        maxIndex := i
        for j := i + 1; j < *m; j++ {
            if T[j].hargaTerjual > T[maxIndex].hargaTerjual {
                maxIndex = j
            }
        }
        T[i], T[maxIndex] = T[maxIndex], T[i]
    }
}


// FUNGSI-FUNGSI MENCARI DATA 
func menuCari(B barang, T transaksi, n,m int){
	var pilih int
	fmt.Println("~~~~~ Menu Mencari Data ~~~~~")
		fmt.Println("1. Cari Data Barang")
		fmt.Println("2. Cari Data Transaksi")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan pilihan (1/2/3): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			menuCariBarang(B,n)
		} else if pilih == 2 {
			clearScreen()
			menuCariTransaksi(T,m)
		} else if pilih == 3 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuCari(B,T,n,m)
	}
}

//CARI DATA BARANG
func menuCariBarang(B barang, n int) {
	var cek int
	var menuCari, cariID int
	var cariNama string
	if n <= 0  {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else {
		fmt.Println("~~ Menu Pencarian Data Barang ~~")
		fmt.Println("1. Cari barang menggunakan ID Barang")
		fmt.Println("2. Cari barang menggunakan Nama Barang")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan Pilihan (1/2/3): ")
		fmt.Scan(&menuCari)
		if menuCari == 1 {
			fmt.Print("Masukkan ID Barang yang dicari: ")
			fmt.Scan(&cariID)
			cek = cariBarangId(B,n,cariID)
			if cek !=-1 {
				clearScreen()
				fmt.Println("Barang dengan ID",cek+1,"ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
				fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				} else {
					clearScreen()
				fmt.Println("Barang tidak dapat ditemukan!\n")
				}
			} else if menuCari == 2{
				fmt.Print("Masukkan Nama Barang yang dicari: ")
				fmt.Scan(&cariNama)
				cek = cariBarangNama(B,n,cariNama)
				if cek !=-1 {
					clearScreen()
					fmt.Println("Barang dengan nama",B[cek].nama,"ditemukan!")
					fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
					fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[cek].idBarang, B[cek].nama, B[cek].kategori, B[cek].hargaBeli, B[cek].hargaJual, B[cek].stock)
				} else {
					clearScreen()
					fmt.Println("Barang tidak dapat ditemukan!\n")
				}
			} else if menuCari == 3{
				clearScreen()
				return
			} else {
				clearScreen()
				fmt.Println("Masukan tidak valid\n")
				menuCariBarang(B,n)
			}
	}
}

func cariBarangId(B barang, n, cariID int ) int {
	urutkanIdBarangAsc(&B,&n)
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if B[mid].idBarang == cariID {
			return mid
		} else if B[mid].idBarang < cariID {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func cariBarangNama(B barang, n int, cariNama string) int {
	for i:=0;i<n;i++{
		if cariNama == B[i].nama {
			return i
		}
	}
	return -1
}

//CARI DATA TRANSAKSI
func menuCariTransaksi(T transaksi, m int){
	var cek, cariID,menuCari int
	if m <= 0  {
		fmt.Println("Tidak ada transaksi yang tersedia!\n")
	} else {
		fmt.Println("~~ Menu Pencarian Data Transaksi ~~")
		fmt.Println("1. Cari barang menggunakan ID Transaksi")
		fmt.Println("2. Kembali")
		fmt.Print("Masukan Pilihan (1/2/3): ")
		fmt.Scan(&menuCari)
		if menuCari == 1 {
			fmt.Print("Masukkan ID Transaksi yang dicari: ")	
			fmt.Scan(&cariID)
			cek = cariTransaksiId(T,m,cariID)
			if cek !=-1 {
				clearScreen()
				fmt.Println("Transaksi dengan ID",cek+1,"ditemukan!")
				fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock Terjual")			
				fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[cek].idTransaksi, T[cek].idBarangTerjual, T[cek].namaTerjual, T[cek].kategoriTerjual, T[cek].hargaTerjual, T[cek].stockTerjual)
			} else {
				clearScreen()
				fmt.Println("Transaksi tidak dapat ditemukan!\n")
			}
		} else if menuCari == 2{
			clearScreen()
			return
		} else {
			clearScreen()
				fmt.Println("Masukan tidak valid\n")
				menuCariTransaksi(T,m)
		}
	}
}

func cariTransaksiId(T transaksi, m, cariID int) int {
	urutkanIdTransAsc(&T,&m)
	low, high := 0, m-1
		for low <= high {
			mid := (low + high) / 2
			if T[mid].idTransaksi == cariID {
				return mid
			} else if T[mid].idTransaksi < cariID {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	return -1
}

//FUNGSI-FUNGSI TAMPIL DATA
func menuTampilData(B barang, T transaksi, n,m int){
	var pilih int
	fmt.Println("~~~~~ Tampil Data ~~~~~")
		fmt.Println("1. Tampilkan Data Barang")
		fmt.Println("2. Tampilkan Data Transaksi")
		fmt.Println("3. Kembali")
		fmt.Print("Masukan pilihan (1/2/3): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			menuTampilDataBarang(&B,&T,&n,&m)
		} else if pilih == 2 {
			clearScreen()
			menuTampilDataTransaksi(&B,&T,&n,&m)
		} else if pilih == 3 {
			clearScreen()
			return
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuTampilData(B,T,n,m)
	}
}

//MENAMPILKAN DATA BARANG
func menuTampilDataBarang(B *barang, T *transaksi, n, m *int){
	var pilih int
	if *n <= 0  {
		fmt.Println("Tidak ada barang yang tersedia!\n")
	} else {
		fmt.Println("~ Menu Tampil Data Barang ~")
		fmt.Println("1. Tampilkan Semua Data Barang")
		fmt.Println("2. Tampilkan 5 Barang Terlaris")
		fmt.Println("3. Kembali")
		fmt.Print("Masukankan pilihan (1/2/3): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			tampilBarang(*B,*n)
		} else if pilih == 2 {
			clearScreen()
			tampilBarangTerlaku(T,m)
		} else if pilih == 3{
			clearScreen()
			menuTampilData(*B,*T,*n,*m)
		}else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuTampilDataBarang(B,T,n,m)
		}
	}
}

func tampilBarang(B barang, n int) {
	if n <=0 {
		fmt.Println("~~~ Data Barang ~~~")
		fmt.Println("Tidak ada barang yang tersedia!")
	} else {
		fmt.Println("~~~ Data Barang ~~~")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Beli", "Harga Jual", "Stock Barang")
		for i := 0; i < n; i++ {
			fmt.Printf("%-20d %-20s %-20s %-20d %-20d %-20d\n", B[i].idBarang, B[i].nama, B[i].kategori, B[i].hargaBeli, B[i].hargaJual, B[i].stock)
		}
	}
	fmt.Println()
	fmt.Println("Sisa modal adalah Rp.",sisaModal,"\n")
}

func tampilBarangTerlaku(T *transaksi, m *int) {
	urutkanStockTerjualDsc(T,m)
	var x int = *m

	if x >= 5{
		x = 5
	}
	fmt.Println("~ 5 Data Barang Terlaku~ ")
	for i:=0; i<x;i++{
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock Terjual")
		fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[i].idTransaksi, T[i].idBarangTerjual, T[i].namaTerjual, T[i].kategoriTerjual, T[i].hargaTerjual, T[i].stockTerjual)
	}
	fmt.Println()
}

//MENAMPILKAN DATA TRANSAKSI
func menuTampilDataTransaksi(B *barang, T *transaksi, n,m *int){
	if *m <=0 {
		fmt.Println("Tidak ada transaksi yang tersedia!\n")
	} else {
		var pilih int
		fmt.Println("~ Menu Tampil Data Transaksi ~")
		fmt.Println("1. Tampilkan Semua Data Transaksi")
		fmt.Println("2. Tampilkan Data Modal")
		fmt.Println("3. Tampilkan Data Pendapatan dan Kerugian")
		fmt.Println("4. Kembali")
		fmt.Print("Masukankan pilihan (1/2/3/4): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			clearScreen()
			tampilDataTransaksi(*T,*m)
		} else if pilih == 2 {
			clearScreen()
			fmt.Println("~~~ Data Modal ~~~")
			fmt.Println("Modal awal sebesar: Rp.",modal)
			fmt.Println("Sisa modal sebesar: Rp.",sisaModal)
			fmt.Println()
		} else if pilih == 3 {
			clearScreen()
			fmt.Println("~~~ Data Pendapatan ~~~")
			fmt.Println("Total keuntungan kotor sebesar: Rp.",keuntungan)
			fmt.Println("Total keuntungan bersih sebesar: Rp.",keuntungan-modal+sisaModal)
			fmt.Println("Total hutang sebesar: Rp.",hutang)
			fmt.Println()
		} else if pilih == 4{
			clearScreen()
			menuTampilData(*B,*T,*n,*m)
		} else {
			clearScreen()
			fmt.Println("Masukan tidak valid!\n")
			menuTampilDataTransaksi(B,T,n,m)
		}
	}
}

func tampilDataTransaksi(T transaksi, m int) {
	if m <= 0 {
		fmt.Println("~~~ Data Transaksi ~~~")
		fmt.Println("Tidak ada transaksi yang tersedia!")
	} else {
		fmt.Println("~~~ Data Transaksi ~~~")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Transaksi", "ID Barang", "Nama Barang", "Kategori Barang", "Harga Jual", "Stock Terjual")
		for i := 0; i < m; i++ {
			fmt.Printf("%-20d %-20d %-20s %-20s %-20d %-20d\n", T[i].idTransaksi, T[i].idBarangTerjual, T[i].namaTerjual, T[i].kategoriTerjual, T[i].hargaTerjual, T[i].stockTerjual)
		}
	}
	fmt.Println()
}

// CLEAR SCREEN
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}