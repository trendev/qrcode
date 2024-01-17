package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func generateQR(u string) {
	err := qrcode.WriteFile(u, qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Printf("impossible to create qrcode: %v\n", err)
	}
}

func main() {
	u := "https://urlz.fr/pj5g"
	fmt.Printf("encoding %q\nlength = %d\n\n", u, len(u))
	generateQR(u)
}
