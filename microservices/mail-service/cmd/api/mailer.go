package main

import (
	"bytes"
	"html/template"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain     string
	Host       string
	Port       int
	Username   string
	Password   string
	Encryption string

	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        interface{}
	DataMap     map[string]interface{}
}

func (m *Mail) SendSMTP(msg *Message) error {
	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]interface{}{
		"Message": msg.Data,
	}

	msg.DataMap = data

	formattedMsg, err := m.buildHTML(msg)
	if err != nil {
		return err
	}

	plainMsg, err := m.buildPlain(msg)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject).SetBody(mail.TextHTML, formattedMsg).AddAlternative(mail.TextPlain, plainMsg)

	if len(msg.Attachments) > 0 {
		for _, att := range msg.Attachments {
			email.AddAttachment(att)
		}
	}

	return email.Send(smtpClient)
}

func (m *Mail) buildHTML(msg *Message) (string, error) {
	templateRender := "./templates/mail.html"

	t, err := template.New("email-html").ParseFiles(templateRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMsg := tpl.String()
	formattedMsg, err = m.inlineCSS(formattedMsg)
	if err != nil {
		return "", err
	}

	return formattedMsg, nil
}

func (m *Mail) buildPlain(msg *Message) (string, error) {
	templateRender := "./templates/mail.plain.html"

	t, err := template.New("email-plain").ParseFiles(templateRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMsg := tpl.String()

	return plainMsg, nil
}

func (m *Mail) inlineCSS(msg string) (string, error) {
	opts := &premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(msg, opts)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *Mail) getEncryption(encryption string) mail.Encryption {
	switch encryption {
	case "ssl":
		return mail.EncryptionSSL
	case "tls":
		return mail.EncryptionTLS
	case "none", "":
		return mail.EncryptionNone

	default:
		return mail.EncryptionSTARTTLS
	}
}
