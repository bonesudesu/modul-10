Nomer 1: 
- Deklarasi Variabel:

Berat: Array dengan kapasitas hingga 1000 elemen bertipe float64 digunakan untuk menyimpan berat badan anak kelinci.
n: Variabel bertipe integer untuk menyimpan jumlah anak kelinci yang datanya akan dimasukkan.

- Input Jumlah Anak Kelinci:

Program meminta pengguna untuk memasukkan jumlah anak kelinci (n) melalui fungsi fmt.Scanln.

- Input Berat Anak Kelinci:

Pengguna diminta untuk memasukkan berat masing-masing anak kelinci satu per satu. Data berat disimpan ke dalam array Berat dengan indeks mulai dari 0 hingga n-1.

- Inisialisasi Nilai Awal untuk Perbandingan:

Variabel minBerat dan maxBerat diinisialisasi dengan nilai berat anak kelinci pertama (Berat[0]) sebagai nilai awal.

- Proses Perbandingan:

Program menggunakan perulangan for untuk membandingkan setiap elemen dalam array Berat:
Jika berat saat ini lebih kecil dari minBerat, maka minBerat diperbarui.
Jika berat saat ini lebih besar dari maxBerat, maka maxBerat diperbarui.

- Output Hasil:

Program mencetak berat terkecil (minBerat) dan berat terbesar (maxBerat) ke layar dengan format dua angka di belakang koma menggunakan fmt.Printf.

Nomer 2 :

- Deklarasi Variabel:

x dan y: Integer untuk jumlah ikan dan jumlah wadah.
beratIkan: Slice bertipe float64 untuk menyimpan berat ikan.
totalBerat: Slice bertipe float64 untuk menyimpan total berat ikan per wadah.
rataRata: Float64 untuk menyimpan rata-rata berat ikan per wadah.

- Pengisian Berat Ikan:

Menggunakan perulangan for untuk memasukkan berat masing-masing ikan ke dalam slice beratIkan.

- Distribusi Berat:

Untuk setiap berat ikan, program menambahkan berat tersebut ke wadah yang sesuai berdasarkan indeks modulo (i%y).

- Perhitungan Rata-Rata Berat:

Program menjumlahkan total berat semua wadah dan membagi hasilnya dengan jumlah wadah (y) untuk mendapatkan rata-rata berat ikan per wadah.

Nomer 3:

- Fungsi hitungMinMax:

Parameter: Slice bertipe float64 yang berisi berat balita.
Proses:
Inisialisasi min dan max dengan elemen pertama array.
Perulangan untuk membandingkan setiap elemen:
Jika elemen lebih kecil dari min, nilai min diperbarui.
Jika elemen lebih besar dari max, nilai max diperbarui.
Keluaran: Nilai minimum (min) dan maksimum (max).

- Fungsi hitungRerata:

Parameter: Slice bertipe float64 yang berisi berat balita.
Proses:
Menjumlahkan semua elemen dalam slice.
Membagi total dengan jumlah elemen slice untuk mendapatkan rata-rata.
Keluaran: Nilai rata-rata berat balita.

- Fungsi main:

Proses Utama:
Meminta input jumlah data berat balita.
Memasukkan data berat balita ke dalam slice arrBalita.
Memanggil fungsi hitungMinMax untuk mendapatkan berat minimum dan maksimum.
Memanggil fungsi hitungRerata untuk mendapatkan rata-rata berat balita.

- Output:
Menampilkan berat minimum, maksimum, dan rata-rata berat balita.







