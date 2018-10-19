package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	send("title", "hello world")
}

func send(title string, body string) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	FROM := os.Getenv("MAILER_USER")
	PASS := os.Getenv("MAILER_PASS")
	TO := "gogawagah@tryzoe.com"

	MSG := "From: " + FROM + "\n" +
		"To: " + TO + "\n" +
		"Subject: " + title + "\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", FROM, PASS, "smtp.gmail.com"),
		FROM, []string{TO}, []byte(MSG))
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Sent")
}
