# <h1 align="center">Laporan Praktikum Modul 05 - Perulangan For</h1>
<p align="center">Zeda Briyan Dinillah - 109092600014</p>

## Dasar Teori

### A. 


### B. 


## Guided

### 1. Nama File: tabungan.go

```go

```
#### Deskripsi
Membuat program yang 

### 2. Nama File: terbesar.go

```go

```
#### Deskripsi
Membuat program yang 

### 3. Nama File: bioskop.go

```go

```
#### Deskripsi
Membuat program yang 

## Unguided

### 1. Prima

```go
package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	jumlahFaktor := 0

	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```

##### Output
![Screenshot Output Unguided](./unguided/prima/unguided_prima_output.png)


#### Deskripsi
Membuat program yang 


### 2. ECCT

```go
package main
import "fmt"

func main() {
	var m int 
	fmt.Scan(&m)
	lulus := 0

	for i := 1; i <= m; i++ {
		var P, V, GA, FS, CI, TC int
		fmt.Scan(&P, &V, &GA, &FS, &CI,&TC)

		total := P + V + GA + FS + CI + TC
		rataRata := float64(total) / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Ya")
			lulus++
		} else {
			fmt.Println("Tidak")
		}
	}
	fmt.Println(lulus)
}
```

##### Output
![Screenshot Output Unguided](./unguided/ecct/unguided_ecct_output.png)

#### Deskripsi
Membuat program 

## Kesimpulan


## Referensi
1. Go Team. (2026). _The Go Programming Language Specification_. Google LLC. Diakses pada 15 Oktober 2026 melalui https://go.dev/ref/spec.
2. Go Team. (2026). _A Tour of Go_. Google LLC. Diakses pada 16 Oktober 2026 melalui https://go.dev/tour/flowcontrol/1.