package monitoring

import (
	"fmt"
	"log"

	c "github.com/Lukpier/flinkmonitoring/config"
	mail "github.com/xhit/go-simple-mail/v2"
)

type MailClient struct {
	config *c.MailConfig
	client *mail.SMTPClient
}

var _ IMailClient = (*MailClient)(nil)

func NewMailClient(config c.MailConfig) *MailClient {

	server := mail.NewSMTPClient()
	server.Host = config.Smtphost
	server.Port = config.Smtpport
	server.Username = config.Sender
	server.Password = config.Password
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	return &MailClient{
		config: &config,
		client: smtpClient,
	}
}

func (mc *MailClient) SendMail(jobId string, body string) error {

	for _, receiver := range mc.config.Receivers {
		// Create email
		email := mail.NewMSG()
		email.SetFrom(mc.config.Sender)
		email.AddTo(receiver)
		email.SetSubject(fmt.Sprintf("Flink Exceptions Report: Job=%s", jobId))
		email.SetBody(mail.TextPlain, body)
		if err := email.Send(mc.client); err != nil {
			return err
		}
	}

	return nil

}
