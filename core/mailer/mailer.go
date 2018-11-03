package mailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"regexp"
)

// Mailer contains context required to send a mail
type Mailer struct {
	from, fromPass string
}

// New returns a new Mailer using parameter email, password
func New(email, password string) (*Mailer, error) {
	if !emailFormat(email) {
		return nil, errors.New("mail: misformatted source email address")
	}

	return &Mailer{email, password}, nil
}

// Send attempts to send a email to the target email using struct information
func (m *Mailer) Send(to, title, body string) error {
	if !emailFormat(to) {
		return errors.New("mail: misformatted target email address")
	}

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", m.from, to, title, body)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", m.from, m.fromPass, "smtp.gmail.com"),
		m.from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

func emailFormat(email string) bool {
	regex := regexp.MustCompile(`^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$`)
	return regex.MatchString(email)
}
