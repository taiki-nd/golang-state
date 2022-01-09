package main

import (
	"io"
	"log"
	"os"
)

func LoggingSetting(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLofFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLofFile)
}

func main() {
	log.Println("logging")
	log.Printf("%T, %v", "test", "test")

	_, err := os.Open("skdfjg")
	if err != nil {
		log.Fatalln("exit", err)
	}

	log.Fatalln("error!")
}
