package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Simulasi Brute Force Sederhana By GunturDeveloper")
	var uPwd string
	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&uPwd)

	pwd := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$&-+,"
	var foundPassword []byte
	startTime := time.Now()

	for _, ch := range uPwd {
		found := false
		for !found {
			guess := pwd[rand.Intn(len(pwd))]
			if byte(guess) == byte(ch) {
				foundPassword = append(foundPassword, byte(guess))
				found = true
				fmt.Printf("Karakter ditemukan: %c\n", guess)
				fmt.Println("Sedang mencari karakter selanjutnya...")
				os.system("clear || cls")
			}
		}
	}

	timeTaken := time.Since(startTime)
	minutes := int(timeTaken.Minutes())
	seconds := int(timeTaken.Seconds()) % 60
	milliseconds := int(timeTaken.Milliseconds()) % 1000

	fmt.Printf("Password ditemukan: %s\n", foundPassword)
	fmt.Printf("Waktu yang dibutuhkan: %d menit %d detik %d milidetik\n", minutes, seconds, milliseconds)
}
