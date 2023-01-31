package main

import (
	"io"
	"log"
	"net"
)

func echo(podkl net.Conn) { // Функция, которая просто echoes полученные данные
	defer podkl.Close()
	buf := make([]byte, 512) // Создаем буфер для хранения полученных данных
	for {
		razm, oshib := podkl.Read(buf[0:]) // Recieve data через conn.Read в буфер.
		if oshib == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if oshib != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Recieved %d bytes:%s\n", razm, string(buf))
		log.Println("Writing data")
		if _, oshib := podkl.Write(buf[0:razm]); oshib != nil { //Отправить инфу через conn.Write
			log.Fatalln("Unable to write a data")
		}
	}
}
func main() {
	slushai, oshibke := net.Listen("tcp", ":20080")
	if oshibke != nil {
		log.Fatalln("Unable to bind a port")
	}
	log.Println("Listening on: 0.0.0.0:20080")
	for {
		konekt, fail := slushai.Accept() //Подождите подключения. Создает объект Conn после подключения.
		log.Println("Recieved connection")
		if fail != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(konekt)
	}
}
