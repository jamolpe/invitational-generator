package mailer

import (
	"errors"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/jamolpe/invitational-generator/internal/parser"
)

type Config struct {
	Server   string
	Port     int
	Email    string
	Password string
}

type MailClient struct {
	To      []string `json:"to" bson:"to"`
	Subject string   `json:"subject" bson:"subject"`
	Body    string   `json:"body" bson:"body"`
}

type Mailer struct {
	config Config
}

func New() Mailer {
	config := configure()
	return Mailer{config}
}

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func configure() Config {
	var config Config
	config.Server = os.Getenv("SERVER")
	config.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	config.Email = os.Getenv("EMAIL")
	config.Password = os.Getenv("PASSWORD")
	return config
}

func (m *Mailer) sendMail(clOptions MailClient) bool {
	body := "To: " + clOptions.To[0] + "\r\nSubject: " + clOptions.Subject + "\r\n" + MIME + "\r\n" + clOptions.Body
	SMTP := fmt.Sprintf("%s:%d", m.config.Server, m.config.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", m.config.Email, m.config.Password, m.config.Server), m.config.Email, clOptions.To, []byte(body)); err != nil {
		return false
	}
	return true
}

func (m *Mailer) Send(templateName string, items interface{}, clOptions MailClient) error {
	parser, err := parser.ParseTemplate(templateName, items)
	if err != nil {
		return err
	}
	clOptions.Body = parser
	if ok := m.sendMail(clOptions); ok {
		log.Printf("Email has been sent to %s\n", clOptions.To)
	} else {
		log.Printf("Failed to send the email to %s\n", clOptions.To)
		return errors.New("failed sending mail")
	}
	return nil
}
