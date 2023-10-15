package loggers

import "log"

func LogInfor(message string) {
	log.Println("Info:", message)
}

func LogError(message string) {
	log.Println("Error:", message)
}
