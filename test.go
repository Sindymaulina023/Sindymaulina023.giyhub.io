package main

import "fmt"

func main() {
    // Deklarasi dan inisialisasi variabel
    var a, b, sum, difference, division, multiplication int

    // Input dua bilangan dari pengguna
    fmt.Print("Masukkan bilangan pertama: ")
    fmt.Scanln(&a)
    fmt.Print("Masukkan bilangan kedua: ")
    fmt.Scanln(&b)

    // Menghitung penjumlahan, pengurangan, pembagian, dan perkalian
    sum = a + b
    difference = a - b
    division = a / b
    multiplication = a * b

    // Menampilkan hasil penjumlahan, pengurangan, pembagian, dan perkalian
    fmt.Printf("Penjumlahan %d dan %d adalah %d\n", a, b, sum)
    fmt.Printf("Pengurangan %d dan %d adalah %d\n", a, b, difference)
    fmt.Printf("Pembagian %d oleh %d adalah %d\n", a, b, division)
    fmt.Printf("Perkalian %d dan %d adalah %d\n", a, b, multiplication)
}