package main

import (
	"log"
	"net"
	"os"

	"github.com/johannesUIA/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.4:8040")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Din melding = ", os.Args[1])
	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("Kryptert melding: ", string(kryptertMelding))
	_, err = conn.Write([]byte(string(kryptertMelding)))

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)
}
