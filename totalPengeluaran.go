func totalPengeluaran(A tabDaftar, N int) {
	var total, i int
	for i = 0; i < N; i++ {
		total += A[i].biaya
	}
	fmt.Printf("Total pengeluaran bulan ini Rp.%d,00\n", total)
}