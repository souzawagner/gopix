package main

import (
	qr "github.com/skip2/go-qrcode"
	"log"
)

func main() {
	var pix = Pix{
		Key:         "73ea37f3-ad17-4a5d-8ee4-f55a5bc5c359",
		Amount:      1.5,
		Description: "Faz um pix ai pia!",
		Merchant: Merchant{
			Name: "Wagner",
			City: "Curitiba",
		},
	}

	var qrcode, _ = pix.QRCode(qr.Medium, 256)

	log.Println("qrcode buffer:", qrcode)
	log.Println("payment code:", pix.String())
	log.Println("creating png")
	pix.File("qrcode.png")
}
