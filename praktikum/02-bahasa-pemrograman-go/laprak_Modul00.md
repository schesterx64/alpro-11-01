# <h1 align="center">Laporan Praktikum Modul 02 - Bahasa Pemrograman Go</h1>
<p align="center">Zeda Briyan Dinillah - 109092600014</p>

## Dasar Teori

### A. Bahasa Pemrograman Go
Go, atau juga sering disebut dengan Golang, merupakan bahasa pemrograman _open-source_ yang dikembangkan oleh Google, diantaranya Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2007, hingga akhirnya dirilis ke publik pada tahun 2009. Go dirancang dengan tujuan performa tinggi namun tetap memiliki sintaks yang ringkas dan mudah dipahami.

### B. Struktur dan Algoritma Bahasa Pemrograman Go

#### 1. `Package main` dan Fungsinya
`Package main` merupakan salah satu nama paket yang ada pada bahasa pemrograman Go. Paket ini berfungsi memberi tahu _compiler_ Go bahwa berkas program dirancang untuk dikompilasi menjadi sebuah aplikasi mandiri yang dapat dieksekusi atau _executable_, bukan hanya sebuah _library_ pendukung. 

#### 2. Tipe Data dan Deklarasi Variabel di Go
Tipe data dalam Go ada beberapa macam, diantaranya yaitu bilangan bulat, bilangan _real_, teks, dan _boolean_. \
a. Bilangan bulat atau _integer_ adalah tipe data yang digunakan untuk menyimpan bilangan bulat, baik positif maupun negatif. Tipe data ini dideklarasikan dengan `int`, `int8`, `int32`, dan `int64`. \
b. Bilangan real atau _float_ adalah tipe data yang digunakan untuk menyimpan bilangan desimal. Tipe data ini dideklarasikan dengan `float32` dan `float 64`. \
c. Teks atau _string_ adalah tipe data yang digunakan untuk menyimpan teks atau karakter. Tipe data ini dideklarasikan dengan tanda petik ganda atau `"..."`. \
d. _Boolean_ adalah tipe data yang hanya memiliki dua kemungkinan, _true_ atau _false_. Tipe data ini dideklarasikan dengan kata kunci `bool`.

Deklarasi variabel di Go dilakukan dengan dua cara utama, yaitu dengan `var` dan `:=`. \
a. Deklarasi `var` digunakan untuk menyebutkan nama variabel dan tipe datanya dengan jelas. \
b. Deklarasi `:=` digunakan agar _compiler_ secara otomatis menebak tipe data berdasarkan nilai didalamnya.

#### 3. Operasi Aritmatika Dasar dan Modulo
Operasi aritmatika dasar dan modulo pada Go adalah operator matematika yang mengolah data angka, baik _integer_ maupun _float_. Operasi aritmatika dasar meliputi penjumlahan (+), pengurangan (-), perkalian (*), dan pembagian (/). Sementara, modulo (%) adalah operator untuk mencari sisa hasil dari pembagian dua bilangan bulat. 

## Guided

### 1. Nama File: skor.go

```go
package main

import "fmt"

func main() {
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung total dan rata-rata (Pembagian bilangan bulat)
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}
```
#### Deskripsi
Membuat program yang membaca _input_ nama seorang siwa, skor matematika, dan skor bahasa Inggris, dan total gabungan dari kedua nilai, lalu mencari rata-rata dari total nilai tadi. Bagian guided yang diimplementasikan yaitu logika dibalik program dan _syntax_ pemrograman. Hasil yang diperoleh adalah _output_ nama siswa, total nilai matematika dan bahasa Inggris, dan rata-rata dari kedua nilai tersebut.

### 2. Nama File: tukar.go

```go
package main

import "fmt"

func main() {
	var a, b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a, b = b, a

	//Output
	fmt.Println(a)
	fmt.Println(b)
}
```
#### Deskripsi
Membuat program yang mampu menukar _input_ a dan b, dan juga sebaliknya. Bagian guided yang diimplementasikan yaitu logika serta _syntax_ pemrograman. Hasil yang diperoleh berupa _output_ yang berkebalikan dari _input_ yang diberikan, misalnya _input_ adalah a dan b, maka _output_ nya adalah b dan a.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. Kalkulator

```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Masukkan angka pertama")
	fmt.Scan(&a)
	fmt.Println("Masukkan angka kedua")
	fmt.Scan(&b)

	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulo := a % b

	fmt.Printf("Tambah = %d, Kurang = %d, Kali = %d, Bagi = %d, Modulo = %d\n", tambah, kurang, kali, bagi, modulo)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)


#### Deskripsi
Membuat program yang menambah, mengurang, mengali, membagi, dan mencari modulo dari dua _input_ yang dimasukkan. Bagian guided yang diimplementasikan adalah cara menggunakan variabel integer, `fmt.Scan`, `fmt.Println`, serta logika matematika dasar. Bagian unguided yang diimplementasikan adalah cara penggunaan _format specifier_ seperti `%d` yang mana digunakan untuk menandai tempat yang akan diisi oleh nilai variabel nantinya, dan `fmt.Printf` untuk mencetak teks dengan format, yang dalam program ini berupa _format specifier_ `%d`. Hasilnya berupa _output_ berupa hasil operasi matematika sederhana dari dua _input_.

### 2. Cacah Uang

```go
package main

import "fmt"

func main() {
	var uang int
	fmt.Println("Masukkan nominal uang")
	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seRibu := sisa / 1000

	fmt.Printf("Sepuluh Ribu = %d, Lima Ribu = %d, Seribu = %d\n", sepuluhRibu, limaRibu, seRibu)
}
```

##### Output
![Screenshot Output Unguided](unguided/[nama_soal]/output.png)

#### Deskripsi
Membuat program yang berfungsi untuk menguraikan uang menjadi nominal pecahan (hanya bilangan bulat) 10.000, 5.000, dan 1.000. Bagian guided yang diimplementasikan yakni cara menggunakan variabel integer, `fmt.Println`, `fmt.Scan`, dan algoritma pembagian dan modulo. Bagian unguided yang diimplementasikan yakni cara menggunakan _format specifier_ seperti `%d`, dan `fmt.Printf`. Hasil yang diperoleh berupa _output_ berapa banyak uang nominal 10.000, 5.000, dan 1.000 dari _input_ total uang.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]

## Referensi
1. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Nama Penulis]. ([Tahun]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
