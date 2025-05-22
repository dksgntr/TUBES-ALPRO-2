func urutLangganan(A *tabDaftar, N int) {
	var i, j int
	var idx, pilih, temp int
	fmt.Println("1. Urut berdasarkan biaya (Selection Sort)")
	fmt.Println("2. Urut berdasarkan tanggal (Insertion Sort)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		for i = 1; i < N; i++ {
			idx = i - 1
			fo j = i j < N j++ {
				if A[idx].biaya <  j].biaya {
					idx  j
				}
			}
			temp = A[i-1]
			A[i-i] = A[idx]
			A[idx] = temp
		}
	} else if pilih == 2 {
		for i = 1; i < N; i++ {
			j = i
			temp = A[i]
			for j > 0 && temp < A[j-1] {
				A[j] = A[j - 1]
				j--
			}
			A[j] = temp
		}
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}