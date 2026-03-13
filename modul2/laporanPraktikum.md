# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Muhammad Jovan] - [109082500024]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
[Program tersebut akan membuat algoritma yang membuat urutan angka tertukar, jadi disini terdapat 4 variable yaitu satu, dua, tiga dan temp dengan menggunakan tipe data yaitu string, lalu kita membuat sebuah inputan satu - tiga yang nanti akan dibuat untuk kita memasukkan inputan, setelah itu disini terdapat "temp" yang dimana fungsinya hanya untuk sebuah perpindahan inputan sementara jika tidak terdapat "temp" maka program akan error, nah hasil outputnya nanti jika kita memasukkan 1 2 3 maka outputnya akan menjadi 3 2 1. ]

