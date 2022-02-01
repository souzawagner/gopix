package main

import (
	"github.com/skip2/go-qrcode"
	"log"
)

func main() {
	var pix = Pix{
		Key:         "73ea37f3-ad17-4a5d-8ee4-f55a5bc5c359",
		Amount:      1.5,
		Description: "Faz um pix ai mano!",
		Merchant: Merchant{
			Name: "Wagner Souza",
			City: "Curitiba",
		},
	}

	var payload = pix.GetPayload()

	log.Println(payload)
	qrcode.WriteFile(payload, qrcode.Medium, 256, "qrcode.png")
}
