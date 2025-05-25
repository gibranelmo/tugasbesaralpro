# Aplikasi Manajemen Stok Bahan Makanan untuk Rumah Tangga

Aplikasi sederhana untuk membantu manajemen stok bahan makanan di rumah tangga. Pengguna dapat mencatat, mencari, mengubah, menghapus, dan melihat daftar stok bahan makanan, serta mendapatkan peringatan kedaluwarsa.

## Fitur Utama

* **Tambah Bahan Makanan:** Menambahkan item baru ke dalam stok dengan nama, jumlah, satuan, dan tanggal kedaluwarsa.
* **Cari Bahan Makanan:** Mencari bahan makanan berdasarkan nama dengan algoritma Binary Search (membutuhkan data terurut berdasarkan nama). Menampilkan informasi lengkap jika ditemukan.
* **Ubah Bahan Makanan:** Mengubah informasi bahan makanan yang sudah ada berdasarkan nama.
* **Hapus Bahan Makanan:** Menghapus bahan makanan dari stok berdasarkan nama.
* **Tampilkan Peringatan Kedaluwarsa:** Menampilkan daftar bahan makanan yang mendekati tanggal kedaluwarsa (7 hari ke depan).
* **Tampilkan Stok:** Menampilkan seluruh daftar stok dan memungkinkan pengurutan berdasarkan nama, tanggal kedaluwarsa, atau jumlah (ascending/descending) menggunakan Selection Sort atau Insertion Sort.
* **Tampilkan Laporan:** Menampilkan laporan sederhana mengenai total jenis bahan makanan yang tersedia.

## Cara Menggunakan

1.  * Pastikan Anda telah menginstal Go (Golang) di sistem Anda.
2.  **Jalankan Aplikasi:**
    * Jalankan perintah: `go run Aplikasi stok bahan makanan.go`
3.  **Interaksi:**
    * Aplikasi akan menampilkan menu utama dengan pilihan angka.
    * Masukkan angka sesuai dengan fitur yang ingin Anda gunakan dan tekan Enter.
    * Ikuti instruksi selanjutnya yang ditampilkan oleh aplikasi.

## Struktur Kode

Aplikasi ini dibangun dengan struktur modular, menggunakan beberapa fungsi (subprogram) untuk setiap fitur utama. Data stok bahan makanan disimpan dalam array statis (`daftarStok`) dengan tipe bentukan `BahanMakanan`.

Algoritma yang diimplementasikan:

* **Pencarian:** Binary Search (membutuhkan data terurut berdasarkan nama).
* **Pengurutan:** Selection Sort dan Insertion Sort.

## Batasan

* Menggunakan array statis dengan ukuran terbatas (`maxStok = 100`).
* Pencarian Binary Search memerlukan data untuk diurutkan berdasarkan nama terlebih dahulu.
* Penyimpanan data tidak persisten (data akan hilang setelah aplikasi ditutup).
* Laporan yang disajikan masih sangat sederhana.

## Pengembangan Lebih Lanjut

Beberapa ide untuk pengembangan aplikasi di masa depan:

* Menggunakan array dinamis (slice) untuk pengelolaan stok yang lebih fleksibel.
* Implementasi penyimpanan data persisten (misalnya, menggunakan file JSON atau database sederhana).
* Fitur pencarian yang lebih canggih (misalnya, berdasarkan kategori atau rentang tanggal kedaluwarsa)..
* Antarmuka Pengguna Grafis (GUI).

## Pembuat

* Muhammad Gibran Elmora Raisha Agnie – 103042310068
* Raefri Arafat Anugrah Natadiredja -  103042400027
