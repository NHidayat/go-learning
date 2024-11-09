/**
DEFER adalah function yang selalu dieksekusi pada akhir sebuah function walaupun terjadi sebuah error
PANIC adalah function untuk memberi error
RECOVER adalah function yg digunakan untuk menangkap data PANIC
*/

package main

import "fmt"

func logging () {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}

	fmt.Println("Selesai running function")
}

func runApplication (value bool) {
	defer logging()
	fmt.Println("Run Application")
	
	if value {
		panic("ERROR!!")
	}

}

func main () {
	runApplication(true)

	fmt.Println("OK")
}