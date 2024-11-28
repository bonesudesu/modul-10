package main

//2311102097 Sya'bananta faqih M lizbar
import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah wadah (y): ")
	fmt.Scan(&x, &y)

	beratIkan := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := range beratIkan {
		fmt.Scan(&beratIkan[i])
	}

	totalBerat := make([]float64, y)
	for i, berat := range beratIkan {
		totalBerat[i%y] += berat
	}

	fmt.Println("Total berat ikan per wadah:")
	for i, total := range totalBerat {
		fmt.Printf("Wadah %d: %.2f\n", i+1, total)
	}

	var rataRata float64
	for _, total := range totalBerat {
		rataRata += total
	}
	fmt.Printf("Rata-rata berat ikan per wadah: %.2f\n", rataRata/float64(y))
}
