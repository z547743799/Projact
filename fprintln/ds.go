package main

import (
	"fmt"
	"os"
)

func main() {
	sd := os.Stdout
	s := 23 - 104
	fmt.Fprintf(sd, "%08b\n", s) // 00100000
	//fmt.Printf("%08b\n", 32)             // 00100000
	//fmt.Print(fmt.Sprintf("%08b\n", 32)) // 00100000
	//fmt.Print(fmt.Errorf("%08b\n", 32))  // 00100000

	fmt.Fprint(os.Stdout, "A")
	fmt.Print("B")
	fmt.Print(fmt.Sprint("C"))
	// ABC

	//fmt.Print("\n")

	fmt.Fprintln(sd, "A") // A
	fmt.Fprintln(sd, "b") // A
	//fmt.Println("B")             // B
	//fmt.Print(fmt.Sprintln("C")) // C
}
