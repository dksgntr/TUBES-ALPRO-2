func hapusLangganan(A *tabDaftar, N *int) {
	var i, j, X int
	fmt.Scan(&X)
	for i = 0; i < *N; i++ {
		if A[i] == X {
			for j = i; j < *N-1; j++ {
				A[j] = A[j+1]
			}
			fmt.Println("Data sudah dihapus")
			return
		}
	}
	*N--
}