package mailer

import (
	"errors"
	"net/smtp"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// Mailer contains context required to send a mail
type Mailer struct {
	toEmail, title, body string
}

// NewMailer returns a new Mailer using parameter target email, title, and body
func NewMailer(email, title, body string) (*Mailer, error) {
	emailFormat := regexp.MustCompile(`^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	if !emailFormat.MatchString(email) {
		return nil, errors.New("mail: misformatted source email address")
	}

	return &Mailer{email, title, body}, nil
}

// Send attempts to send a email to the target email using struct information
func (m *Mailer) Send() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	FROM := os.Getenv("MAILER_USER")
	PASS := os.Getenv("MAILER_PASS")
	TO := m.toEmail

	MSG := "From: " + FROM + "\n" +
		"To: " + TO + "\n" +
		"Subject: " + m.title + "\n\n" +
		m.body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", FROM, PASS, "smtp.gmail.com"),
		FROM, []string{TO}, []byte(MSG))
	if err != nil {
		return err
	}

	return nil
}