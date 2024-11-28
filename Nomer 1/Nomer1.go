package main

//2311102097 Sya'bananta faqih M lizbar
import "fmt"

func main() {

	var Berat [1000]float64

	var n int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scanln(&n)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Scan(&Berat[i])
	}

	minBerat := Berat[0]
	maxBerat := Berat[0]

	for i := 1; i < n; i++ {
		if Berat[i] < minBerat {
			minBerat = Berat[i]
		}
		if Berat[i] > maxBerat {
			maxBerat = Berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
