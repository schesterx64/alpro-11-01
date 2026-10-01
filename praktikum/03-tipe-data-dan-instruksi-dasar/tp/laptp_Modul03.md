# <h1 align="center">Tugas Pendahuluan Modul 03 - Variabel dan Operator Bahasa Pemrograman Go</h1>
<p align="center">Zeda Briyan Dinillah - 109092600014</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	
	var x, y int
	fmt.Println("Jumlah Kue")
	fmt.Scan(&y)
	fmt.Println("Jumlah Anggota Keluarga")
	fmt.Scan(&x)

	modulo := y % x

	fmt.Printf("Sisa = %d", modulo)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/sisa/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

### 2. Konversi

```go
package main

import "fmt"

func main() {
	var x float64

	fmt.Println("Masukkan jarak dalam mil")
	fmt.Scan(&x)

	kilometer := x * 1.6

	fmt.Println("Konversi ke kilometer", kilometer)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/konversi/output.png)

### 3. Boolean

```go
package main

import "fmt"

func main() {
	var bool bool

	fmt.Println("Masukkan niai Boolean")
	fmt.Scan(&bool)

	fmt.Println(bool)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instriksi-dasar/tp/sisa/output.png)


#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

#### Deskripsi
[Tuliskan ringkasan proses praktikum: apa yang dikerjakan, bagian guided dan unguided yang diimplementasikan, serta hasil yang diperoleh.]

## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]