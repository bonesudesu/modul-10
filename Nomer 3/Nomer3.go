package main

//2311102097 Sya'bananta faqih M lizbar
import "fmt"

func hitungMinMax(arr []float64) (min, max float64) {
	min, max = arr[0], arr[0]
	for _, berat := range arr {
		if berat < min {
			min = berat
		}
		if berat > max {
			max = berat
		}
	}
	return
}

func hitungRerata(arr []float64) (total float64) {
	for _, berat := range arr {
		total += berat
	}
	return total / float64(len(arr))
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	arrBalita := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBalita[i])
	}

	bMin, bMax := hitungMinMax(arrBalita)
	rerata := hitungRerata(arrBalita)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
