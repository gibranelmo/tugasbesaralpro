package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const maxStok = 100

type BahanMakanan struct {
	Nama          string
	Jumlah        int
	Satuan        string
	TanggalKedaluwarsa time.Time
}

var daftarStok [maxStok]BahanMakanan
var jumlahStok int = 0

func tambahBahanMakanan(nama string, jumlah int, satuan string, tanggalKedaluwarsa string) {
	if jumlahStok < maxStok {
		tanggal, err := time.Parse("2006-01-02", tanggalKedaluwarsa)
		if err != nil {
			fmt.Println("Format tanggal tidak valid (YYYY-MM-DD).")
			return
		}
		daftarStok[jumlahStok] = BahanMakanan{Nama: nama, Jumlah: jumlah, Satuan: satuan, TanggalKedaluwarsa: tanggal}
		jumlahStok++
		fmt.Println("Bahan makanan berhasil ditambahkan.")
	} else {
		fmt.Println("Stok penuh. Tidak dapat menambahkan bahan makanan lagi.")
	}
}

func binarySearch(nama string, data []BahanMakanan) int {
	low := 0
	high := len(data) - 1
	for low <= high {
		mid := (low + high) / 2
		if strings.ToLower(data[mid].Nama) == strings.ToLower(nama) {
			return mid
		} else if strings.ToLower(data[mid].Nama) < strings.ToLower(nama) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func ubahBahanMakanan(namaCari string, namaBaru string, jumlahBaru int, satuanBaru string, tanggalKedaluwarsaBaru string) {
	index := binarySearch(namaCari, daftarStok[:jumlahStok])
	if index != -1 {
		tanggal, err := time.Parse("2006-01-02", tanggalKedaluwarsaBaru)
		if err != nil {
			fmt.Println("Format tanggal tidak valid (YYYY-MM-DD).")
			return
		}
		daftarStok[index].Nama = namaBaru
		daftarStok[index].Jumlah = jumlahBaru
		daftarStok[index].Satuan = satuanBaru
		daftarStok[index].TanggalKedaluwarsa = tanggal
		fmt.Printf("Data bahan makanan '%s' berhasil diubah.\n", namaCari)
	} else {
		fmt.Printf("Bahan makanan '%s' tidak ditemukan.\n", namaCari)
	}
}

func hapusBahanMakanan(namaHapus string) {
	index := binarySearch(namaHapus, daftarStok[:jumlahStok])
	if index != -1 {
		for i := index; i < jumlahStok-1; i++ {
			daftarStok[i] = daftarStok[i+1]
		}
		jumlahStok--
		fmt.Printf("Bahan makanan '%s' berhasil dihapus.\n", namaHapus)
	} else {
		fmt.Printf("Bahan makanan '%s' tidak ditemukan.\n", namaHapus)
	}
}

func peringatanKedaluwarsa() {
	fmt.Println("\n--- Peringatan Bahan Makanan Mendekati Kedaluwarsa ---")
	now := time.Now()
	for i := 0; i < jumlahStok; i++ {
		selisih := daftarStok[i].TanggalKedaluwarsa.Sub(now).Hours() / 24
		if selisih <= 7 {
			fmt.Printf("- %s akan kedaluwarsa pada: %s (tersisa %d hari)\n",
				daftarStok[i].Nama, daftarStok[i].TanggalKedaluwarsa.Format("2006-01-02"), int(selisih))
		}
	}
	if jumlahStok == 0 {
		fmt.Println("Stok kosong.")
	}
}

func selectionSort(kategori string, urutan string, data []BahanMakanan) []BahanMakanan {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			switch kategori {
			case "nama":
				if (urutan == "asc" && strings.ToLower(data[j].Nama) < strings.ToLower(data[minIndex].Nama)) ||
					(urutan == "desc" && strings.ToLower(data[j].Nama) > strings.ToLower(data[minIndex].Nama)) {
					minIndex = j
				}
			case "tanggal":
				if (urutan == "asc" && data[j].TanggalKedaluwarsa.Before(data[minIndex].TanggalKedaluwarsa)) ||
					(urutan == "desc" && data[j].TanggalKedaluwarsa.After(data[minIndex].TanggalKedaluwarsa)) {
					minIndex = j
				}
			case "jumlah":
				if (urutan == "asc" && data[j].Jumlah < data[minIndex].Jumlah) ||
					(urutan == "desc" && data[j].Jumlah > data[minIndex].Jumlah) {
					minIndex = j
				}
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}

func insertionSort(kategori string, urutan string, data []BahanMakanan) []BahanMakanan {
	n := len(data)
	for i := 1; i < n; i++ {
		key := data[i]
		j := i - 1
		for j >= 0 {
			compare := false
			switch kategori {
			case "nama":
				if (urutan == "asc" && strings.ToLower(data[j].Nama) > strings.ToLower(key.Nama)) ||
					(urutan == "desc" && strings.ToLower(data[j].Nama) < strings.ToLower(key.Nama)) {
					compare = true
				}
			case "tanggal":
				if (urutan == "asc" && data[j].TanggalKedaluwarsa.After(key.TanggalKedaluwarsa)) ||
					(urutan == "desc" && data[j].TanggalKedaluwarsa.Before(key.TanggalKedaluwarsa)) {
					compare = true
				}
			case "jumlah":
				if (urutan == "asc" && data[j].Jumlah > key.Jumlah) ||
					(urutan == "desc" && data[j].Jumlah < key.Jumlah) {
					compare = true
				}
			}
			if compare {
				data[j+1] = data[j]
				j--
			} else {
				break
			}
		}
		data[j+1] = key
	}
	return data
}

func tampilkanStok(urutBerdasarkan string, urutan string, metodeUrut string) {
	fmt.Println("\n--- Daftar Stok Bahan Makanan ---")
	if jumlahStok > 0 {
		var sortedStok []BahanMakanan
		copy(sortedStok, daftarStok[:jumlahStok])

		if metodeUrut == "selection" {
			sortedStok = selectionSort(urutBerdasarkan, urutan, daftarStok[:jumlahStok])
		} else if metodeUrut == "insertion" {
			sortedStok = insertionSort(urutBerdasarkan, urutan, daftarStok[:jumlahStok])
		} else {
			sortedStok = daftarStok[:jumlahStok]
		}

		for i := 0; i < len(sortedStok); i++ {
			fmt.Printf("%d. Nama: %s, Jumlah: %d %s, Kedaluwarsa: %s\n",
				i+1, sortedStok[i].Nama, sortedStok[i].Jumlah, sortedStok[i].Satuan, sortedStok[i].TanggalKedaluwarsa.Format("2006-01-02"))
		}
	} else {
		fmt.Println("Stok kosong.")
	}
}

func tampilkanLaporan() {
	fmt.Println("\n--- Laporan Stok ---")
	fmt.Printf("Total Bahan Makanan Tersedia: %d jenis\n", jumlahStok)
}

func main() {
	for {
		fmt.Println("\n--- Aplikasi Manajemen Stok Bahan Makanan ---")
		fmt.Println("1. Tambah Bahan Makanan")
		fmt.Println("2. Cari Bahan Makanan")
		fmt.Println("3. Ubah Bahan Makanan")
		fmt.Println("4. Hapus Bahan Makanan")
		fmt.Println("5. Tampilkan Peringatan Kedaluwarsa")
		fmt.Println("6. Tampilkan Stok (Urutkan)")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan string
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			fmt.Print("Nama Bahan Makanan: ")
			var nama string
			fmt.Scanln(&nama)
			fmt.Print("Jumlah: ")
			var jumlah int
			fmt.Scanln(&jumlah)
			fmt.Print("Satuan: ")
			var satuan string
			fmt.Scanln(&satuan)
			fmt.Print("Tanggal Kedaluwarsa (YYYY-MM-DD): ")
			var tanggalKedaluwarsa string
			fmt.Scanln(&tanggalKedaluwarsa)
			tambahBahanMakanan(nama, jumlah, satuan, tanggalKedaluwarsa)
		case "2":
			tempStok := make([]BahanMakanan, jumlahStok)
			copy(tempStok, daftarStok[:jumlahStok])
			tempStok = selectionSort("nama", "asc", tempStok)

			fmt.Print("Masukkan nama bahan makanan yang dicari: ")
			var cariNamaBin string
			fmt.Scanln(&cariNamaBin)

			index := binarySearch(cariNamaBin, tempStok)
			if index != -1 {
				fmt.Println("\n--- Informasi Bahan Makanan Ditemukan ---")
				fmt.Printf("Nama: %s\n", tempStok[index].Nama)
				fmt.Printf("Jumlah: %d %s\n", tempStok[index].Jumlah, tempStok[index].Satuan)
				fmt.Printf("Tanggal Kedaluwarsa: %s\n", tempStok[index].TanggalKedaluwarsa.Format("2006-01-02"))
			} else {
				fmt.Printf("Bahan makanan '%s' tidak ditemukan.\n", cariNamaBin)
			}
		case "3":
			fmt.Print("Masukkan nama bahan makanan yang ingin diubah: ")
			var namaUbah string
			fmt.Scanln(&namaUbah)
			fmt.Print("Nama Bahan Makanan Baru: ")
			var namaBaru string
			fmt.Scanln(&namaBaru)
			fmt.Print("Jumlah Baru: ")
			var jumlahBaru int
			fmt.Scanln(&jumlahBaru)
			fmt.Print("Satuan Baru: ")
			var satuanBaru string
			fmt.Scanln(&satuanBaru)
			fmt.Print("Tanggal Kedaluwarsa Baru (YYYY-MM-DD): ")
			var tanggalKedaluwarsaBaru string
			fmt.Scanln(&tanggalKedaluwarsaBaru)
			ubahBahanMakanan(namaUbah, namaBaru, jumlahBaru, satuanBaru, tanggalKedaluwarsaBaru)
		case "4":
			fmt.Print("Masukkan nama bahan makanan yang ingin dihapus: ")
			var namaHapus string
			fmt.Scanln(&namaHapus)
			hapusBahanMakanan(namaHapus)
		case "5":
			peringatanKedaluwarsa()
		case "6":
			if jumlahStok > 0 {
				fmt.Print("Urutkan berdasarkan (nama/tanggal/jumlah): ")
				var urutBerdasarkan string
				fmt.Scanln(&urutBerdasarkan)
				urutBerdasarkan = strings.ToLower(urutBerdasarkan)
				if urutBerdasarkan == "nama" || urutBerdasarkan == "tanggal" || urutBerdasarkan == "jumlah" {
					fmt.Print("Urutan (asc/desc): ")
					var urutan string
					fmt.Scanln(&urutan)
					urutan = strings.ToLower(urutan)
					if urutan == "asc" || urutan == "desc" {
						fmt.Print("Metode Urut (selection/insertion): ")
						var metodeUrut string
						fmt.Scanln(&metodeUrut)
						metodeUrut = strings.ToLower(metodeUrut)
						if metodeUrut == "selection" || metodeUrut == "insertion" {
							tampilkanStok(urutBerdasarkan, urutan, metodeUrut)
						} else {
							fmt.Println("Metode urut tidak valid.")
						}
					} else {
						fmt.Println("Urutan tidak valid. Gunakan 'asc' atau 'desc'.")
					}
				} else {
					fmt.Println("Kategori pengurutan tidak valid. Gunakan 'nama', 'tanggal', atau 'jumlah'.")
				}
			} else {
				fmt.Println("Stok kosong.")
			}
		case "7":
			tampilkanLaporan()
		case "8":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}