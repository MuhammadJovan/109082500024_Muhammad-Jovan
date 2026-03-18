# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Muhammad Jovan] - [109082500024]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i<= n; i++ {
		hasil *= i
	}
	return hasil
}

	func permutasi(n,r int) int {
	return faktorial(n) / faktorial(n-r)
}

	func kombinasi(n,r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main () {
	var a,b,c,d int
	fmt.Scan(&a,&b,&c,&d)

	fmt.Println(permutasi(a,c), kombinasi(a,c))
	fmt.Println(permutasi(b,d), kombinasi(b,d))
}
```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal1.png)
[Penjelasan]

//

### 2. [Soal]
#### soal2.go

```go
package main

import "fmt"

func f(x int) int {
	return	 x * x
}

func g(x int) int {
	return	x - 2
}

func h(x int) int {
	return	x + 1
}

func main() {
	var a,b,c int
	fmt.Scan(&a,&b,&c)

	hasil := f(g(h(a)))
	hasil2 := g(h(f(b)))
	hasil3 := h(f(g(c)))

	fmt.Println(hasil)
	fmt.Println(hasil2)
	fmt.Println(hasil3)
}

```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal2.png)
[Penjelasan] 

//

### 3. [Soal]
#### soal2.go

```go
package main

import (
		"fmt"
		"math"
)

func jarak(a,b,c,d float64) float64 {
	return math.Sqrt((a-c)*(a-c)+(b-d)*(b-d))
}

func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	dalam1 := diDalam(cx1, cy1, r1, x, y)
	dalam2 := diDalam(cx2, cy2, r2, x, y)

	if  dalam1 && dalam2 {
		fmt.Println("titik di dalam lingkaran 1 dan 2")
	} else if dalam1{
		fmt.Println("titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("titik di dalam lingkaran 2")
	} else {
		fmt.Println("titik di luar lingkaran 1 dan 2")
	}
}


```
### Output Unguided :	

##### Output 
![Screenshot Output Unguided 1_](https://github.com/MuhammadJovan/109082500024_Muhammad-Jovan/blob/main/modul2/output/output-soal3.png)
[Penjelasan]