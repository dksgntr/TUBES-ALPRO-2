func cariLangganan(A tabDaftar, N int) int {
	var i, ketemu, X int
	ketemu = -1
	fmt.Scan(&X)
	for i = 0; i < N; i++ {
		if A[i] = X {
			ketemu = i
		}
	}
	return ketemu
}