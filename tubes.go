package main

import "fmt"

func main() {
	var berat float64
	var layanan int

	// Tarif
	tarifPerKg := 5000.0
	tarifPerKm := 2000.0

	// ARRAY KOTA & NILAI POSISI
	kota := []string{
		"Jakarta",
		"Karawang",
		"Bandung",
		"Ngawi",
		"Banyuwangi",
	}

	posisi := []float64{
		0,    // Jakarta
		50,   // Karawang
		150,  // Bandung
		575,  // Ngawi
		1140, // Banyuwangi
	}

	fmt.Println("KALKULATOR BIAYA PENGIRIMAN BARANG")

	// Input berat

	fmt.Print("Masukkan berat barang (kg): ")
	fmt.Scan(&berat)

	for berat <= 0 {
		fmt.Println("ERROR: Berat harus > 0")
		fmt.Print("Masukkan berat barang (kg): ")
		fmt.Scan(&berat)
	}

	// Pilih kota asal
	fmt.Println("\nPilih Kota Asal:")
	for i, k := range kota {
		fmt.Printf("%d. %s\n", i+1, k)
	}

	var KotaAsal int
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&KotaAsal)

	for KotaAsal < 1 || KotaAsal > len(kota) {
		fmt.Println("ERROR: Kota asal tidak valid")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&KotaAsal)
	}
	KotaAsal--

	// Pilih kota tujuan
	fmt.Println("\nPilih Kota Tujuan:")
	for i, k := range kota {
		fmt.Printf("%d. %s\n", i+1, k)
	}

	var KotaTujuan int

	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&KotaTujuan)

	for KotaTujuan < 1 || KotaTujuan > len(kota) || KotaAsal == KotaTujuan-1 {
		fmt.Println("ERROR: Kota tujuan tidak valid atau sama dengan kota asal")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&KotaTujuan)

	}
	KotaTujuan--

	// Hitung Jarak
	jarak := posisi[KotaAsal] - posisi[KotaTujuan]
	if jarak < 0 {
		jarak = -jarak
	}

	// Pilih layanan
	fmt.Println("\nPilih layanan:")
	fmt.Println("1. Reguler")
	fmt.Println("2. Express")
	fmt.Println("3. Cargo")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&layanan)

	biayaDasar := (berat * tarifPerKg) + (jarak * tarifPerKm)

	var faktor, estimasiHari float64

	for layanan < 1 || layanan > 3 || (layanan == 3 && berat < 5) {
		if layanan == 3 && berat < 5 {
			fmt.Println("ERROR: Cargo minimal 5 kg")
		} else {
			fmt.Println("ERROR: Layanan tidak valid")
		}

		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&layanan)
	}

	switch layanan {
	case 1:
		faktor = 1.0
		estimasiHari = jarak / 300
	case 2:
		faktor = 1.5
		estimasiHari = jarak / 500
	case 3:
		faktor = 0.8
		estimasiHari = jarak / 200
	}

	total := biayaDasar * faktor

	// Estimasi waktu
	var estimasiOutput string
	if estimasiHari < 1 {
		estimasiOutput = fmt.Sprintf("%.1f Jam", estimasiHari*24)
	} else {
		estimasiOutput = fmt.Sprintf("%.1f Hari", estimasiHari)
	}

	// Output
	fmt.Println("\nRINCIAN PENGIRIMAN")
	fmt.Printf("Asal          : %s\n", kota[KotaAsal])
	fmt.Printf("Tujuan        : %s\n", kota[KotaTujuan])
	fmt.Printf("Jarak         : %.0f km\n", jarak)
	fmt.Printf("Berat         : %.1f kg\n", berat)
	fmt.Printf("Total Biaya   : Rp %.0f\n", total)
	fmt.Printf("Estimasi Waktu: %s\n", estimasiOutput)
}
