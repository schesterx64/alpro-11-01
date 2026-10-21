package main
import "fmt"

func main() {
	var PIN, pilih, saldo, tarik int
	saldo = 10000000
	i := 0
	benar := false

	for i < 3 {
		fmt.Println("Masukkan PIN ATM Anda:")
		fmt.Scan(&PIN)
		i++

		if PIN == 1234 {
			fmt.Println("Login berhasil!")
			benar = true
			break
		} else {
			if i < 3 {
				fmt.Println("PIN salah. Silakan coba lagi.")
			}
		}
	}

	if benar == false {
		fmt.Println("ATM terblokir. Program selesai.")
		return

	}
	
	menu:
	for {
		fmt.Println("Menu ATM")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			fmt.Println("Saldo Anda:", saldo)
		case 2:
			fmt.Println("Masukkan jumlah uang yang ingin ditarik:")
			fmt.Scan(&tarik)
			saldo -= tarik
			fmt.Println("Anda berhasil menarik uang sebesar Rp.", tarik)
			fmt.Println("Saldo tersisa:", saldo)
		case 3:
			fmt.Println("Keluar dari sistem. Terima kasih!")
			break menu
		}
	}
}