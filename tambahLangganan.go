func tambahLangganan(A *tabDaftar, n *int) {
	if n >= NMAX {
		fmt.Println("Data kamu penuh:)")
		return
	}

	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[n].nama)
	fmt.Print("Kategori: ")
	fmt.Scan(&A[n].kategori)
	fmt.Print("Biaya Bulanan: ")
	fmt.Scan(&A[n].biaya)
	fmt.Print("Tanggal Pembayaran (1-31): ")
	fmt.Scan(&A[n].tglBayar)
	for A[n].tglBayar < 1 && A[n].tglBayar > 31 {
		fmt.Print("Coba lagi, tanggal yang kamu masukin ga valid\nTanggal Pembayaran (1-31): ")
		fmt.Scan(&A[n].tglBayar)
	}
	fmt.Print("Metode Pembayaran: ")
	fmt.Scan(&A[n].metodBayar)
	A[n].status = true
	n++

	fmt.Println("Data berhasil ditambah:)")
}