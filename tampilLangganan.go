func tampilLangganan(A tabDaftar, n int) {
	var i int
	if n == 0 {
		fmt.Println("Kamu belum punya daftar langganan :)")
		return
	}

	for i = 0; i < n; i++ {
		fmt.Printf("%d %s(%s) Rp.%d,00 tanggal bayar %d (%s)\n", i+1, A[i].nama, A[i].kategori, A[i].biaya, A[i].tglBayar, A[i].status)
	}
}