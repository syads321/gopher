package helper

import (
	"log"
	"os"
)

// LogInfo dfjsdkjf
func LogInfo(msg string) {
	f, err := os.OpenFile("log/text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "GOPHER - ", log.LstdFlags)
	logger.Println(msg)
}

func LogFatal(msg string) {
	f, err := os.OpenFile("log/text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "GOPHER - ", log.LstdFlags)
	logger.Fatal(msg)
}
