package main

import "fmt"

const NMAX = 100

type daftar struct {
	nama, kategori, metodBayar string
	biaya, tglBayar            int
	status                     bool
}

var tabDaftar [NMAX]daftar

func main() {
	var data tabDaftar
	var nData, pilih int
	var idxCL int
	nData = 0
	for {
		menu()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tampilLangganan(data, nData)
		case 2:
			tambahLangganan(&data, &nData)
		case 3:
			editLangganan(&data, &nData)
		case 4:
			hapusLangganan(&data, &nData)
		case 5:
			idxCL = cariLangganan(data, nData)
			if cariLangganan == -1 {
				fmt.Printlh("Data yang kamu cari ga ketemu nih")
			} else {
				fmt.Println("Data yang kamu cari ada di urutan ke : ", idxCL)
			}
		case 6:
			urutLangganan(&data, nData)
		case 7:
			totalPengeluaran(data, nData)
		case 8:
			rekomendasiHemat()
		case 9:
			fmt.Println("Terimakasih :)")
			return
		default:
			fmt.Println("Pilihan kamu ga valid nih")
		}

	}
}

func menu() {
	fmt.Println("\n\t====== MANAJEMEN SUBS & KEUANGAN ======")
	fmt.Println("1. Lihat Semua Langganan")
	fmt.Println("2. Tambah Langganan")
	fmt.Println("3. Edit Langganan")
	fmt.Println("4. Hapus Langganan")
	fmt.Println("5. Cari Langganan")
	fmt.Println("6. Urutkan Langganan")
	fmt.Println("7. Total Pengeluaran Bulan Ini")
	fmt.Println("8. Rekomendasi Penghematan")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu: ")
}
