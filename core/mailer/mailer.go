package mailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"regexp"
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
	from := os.Getenv("MAILER_USER")
	pass := os.Getenv("MAILER_PASS")
	to := m.toEmail

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", from, to, m.title, m.body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
