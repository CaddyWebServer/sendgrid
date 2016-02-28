// Package sendgrid implements the emailsender driver interface from
// github.com/itsabot/abot/shared/interface/emailsender
package sendgrid

import (
	"github.com/itsabot/abot/shared/interface/emailsender"
	"github.com/itsabot/abot/shared/interface/emailsender/driver"
	"github.com/sendgrid/sendgrid-go"
)

type drv struct{}

func (d *drv) Open(name string) (driver.Conn, error) {
	c := conn{sendgrid.NewSendGridClientWithApiKey(name)}
	return &c, nil
}

func init() {
	emailsender.Register("sendgrid", &drv{})
}

type conn struct {
	*sendgrid.SGClient
}

func (c *conn) SendHTML(to []string, from, subj, html string) error {
	msg := sendgrid.NewMail()
	for _, t := range to {
		msg.AddTo(t)
	}
	msg.SetSubject(subj)
	msg.SetFrom(from)
	msg.SetHTML(html)
	return c.Send(msg)
}

func (c *conn) SendPlainText(to []string, from, subj, plaintext string) error {
	msg := sendgrid.NewMail()
	for _, t := range to {
		msg.AddTo(t)
	}
	msg.SetSubject(subj)
	msg.SetFrom(from)
	msg.SetText(plaintext)
	return c.Send(msg)
}

func (c *conn) Close() error {
	return nil
}
