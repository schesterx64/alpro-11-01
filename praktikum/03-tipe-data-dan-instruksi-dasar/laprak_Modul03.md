# <h1 align="center">Laporan Praktikum Modul 03 - Variable dan Operator</h1>
<p align="center">Zeda Briyan Dinillah - 109092600014</p>

## Dasar Teori

### A. 


### B. 

#### 1. 


#### 2. 


#### 3. 


## Guided

### 1. Nama File: kasir.go

```go
package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan nominal: ")
	fmt.Scan(&x)

	var sepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println(sepuluhRibuan, limaRibuan, seribuan)
}
```
#### Deskripsi
Membuat program yang memecah nominal _input_ menjadi pecahan sepuluh ribu, lima ribu, dan seribu. Bagian guided yang diimplementasikan yaitu operasi pembagian dan modulo. Hasil yang diperoleh berupa jumlah uang sepuluh ribu, lima ribu, dan seribu dari input awal.

### 2. Nama File: konversi.go

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukkan suhu: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}
```
#### Deskripsi
Membuat program yang melakukan konversi suhu dari derajar Celcius ke derajat Kelvin. Bagian guided yang diimplementasikan yaitu operasi penjumlahan dari _input_ dan nilai konversi, dan cara menggunakan tipe data `float64`. Hasil yang diperoleh berupa konversi dari derajat Celcius ke derajat Kelvin.

### 3. Nama File: tukar.go

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}
```
#### Deskripsi
Membuat program yang menukar _input_ a, b, c menjadi c, a, b. Bagian guided yang diimplementasikan yaitu cara menukar variabel dan cara menggunakan `temp`. Hasil yang diperoleh berupa tiga angka dari _input_ awal yang telah ditukar posisinya.

## Unguided

### 1. Konversi Suhu

```go
package main 

import "fmt"

func main() {
	fmt.Println("Masukkan suhu dalam Celcius:")
	var C float64
	fmt.Scan(&C)

	fmt.Println("Konversi ke Reamur:", (4.0/5.0) * C)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/schesterx64/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversiSuhu/unguided_konversiSuhu_output.png)


#### Deskripsi
Membuat program yang melakukan konversi suhu dari derajat Celcius ke derajat Reamur. Bagian guided yang diimplementasikan yaitu operasi penjumlahan dari _input_ dan nilai konversi, dan cara menggunakan tipe data `float64`. Hasil yang diperoleh berupa konversi dari derajat Celcius ke derajat Reamur dengan menggunakan rumus $R = (4 / 5) * C$.


### 2. Konversi Hari

```go
package main

import "fmt"

func main() {
	var hari, sisaHari, minggu, bulan, tahun int
	fmt.Println("Masukkan jumlah hari:")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari = hari % 360
	bulan = hari / 30
	hari = hari % 30
	minggu = hari / 7
	sisaHari = hari % 7
	fmt.Printf("%d tahun, %d bulan, %d minggu, %d hari\n", tahun, bulan, minggu, sisaHari)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/schesterx64/alpro-11-01/blob/main/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversiHari/unguided_konversiHari_output.png)

#### Deskripsi
Membuat program yang melakukan konversi jumlah hari menjadi tahun, bulan, minggu, dan sisa hari. Bagian guided yang diimplementasikan operasi pembagian dan modulo. Bagian unguided yang diimplementasikan yaitu penggunaan _input-specifier_ `%d`. Hasil yang diperoleh berupa jumlah tahun, bulan, minggu, dan hari yang tersisa dari _input_ awal menggunakan operasi pembagian dan modulo.


## Kesimpulan
Berdasarkan pembahasan teori dan praktikum yang telah dilakukan, dapat ditarik kesimpulan bahwa:




## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 24 September 2026 melalui https://go.dev/ref/spec.
2. Go Team. (2026). _Standard Library Documentation_. Google LLC. Diakses pada 24 September 2026 melalui https://pkg.go.dev/fmt#section-documentation.