package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"

	"github.com/goravel/framework/contracts/mail"
	contractqueue "github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
)

type Email struct {
	clone    int
	content  mail.Content
	from     mail.From
	to       []string
	cc       []string
	bcc      []string
	attaches []string
}

func NewEmail() mail.Mail {
	return &Email{}
}

func (r *Email) Content(content mail.Content) mail.Mail {
	instance := r.instance()
	instance.content = content

	return instance
}

func (r *Email) From(from mail.From) mail.Mail {
	instance := r.instance()
	instance.from = from

	return instance
}

func (r *Email) To(to []string) mail.Mail {
	instance := r.instance()
	instance.to = to

	return instance
}

func (r *Email) Cc(cc []string) mail.Mail {
	instance := r.instance()
	instance.cc = cc

	return instance
}

func (r *Email) Bcc(bcc []string) mail.Mail {
	instance := r.instance()
	instance.bcc = bcc

	return instance
}

func (r *Email) Attach(files []string) mail.Mail {
	instance := r.instance()
	instance.attaches = files

	return instance
}

func (r *Email) Send() error {
	return SendMail(r.content.Subject, r.content.Html, r.from.Address, r.from.Name, r.to, r.cc, r.bcc, r.attaches)
}

func (r *Email) Queue(queue *mail.Queue) error {
	job := facades.Queue.Job(&SendMailJob{}, []contractqueue.Arg{
		{Value: r.content.Subject, Type: "string"},
		{Value: r.content.Html, Type: "string"},
		{Value: r.from.Address, Type: "string"},
		{Value: r.from.Name, Type: "string"},
		{Value: r.to, Type: "[]string"},
		{Value: r.cc, Type: "[]string"},
		{Value: r.bcc, Type: "[]string"},
		{Value: r.attaches, Type: "[]string"},
	})
	if queue != nil {
		if queue.Connection != "" {
			job.OnConnection(queue.Connection)
		}
		if queue.Queue != "" {
			job.OnQueue(queue.Queue)
		}
	}

	return job.Dispatch()
}

func (r *Email) instance() *Email {
	if r.clone == 0 {
		return &Email{clone: 1}
	}

	return r
}

func SendMail(subject, html string, fromAddress, fromName string, to, cc, bcc, attaches []string) error {
	e := email.NewEmail()
	if fromAddress == "" {
		e.From = fmt.Sprintf("%s <%s>", facades.Config.GetString("mail.from.name"), facades.Config.GetString("mail.from.address"))
	} else {
		e.From = fmt.Sprintf("%s <%s>", fromName, fromAddress)
	}

	e.To = to
	e.Bcc = bcc
	e.Cc = cc
	e.Subject = subject
	e.HTML = []byte(html)

	for _, attach := range attaches {
		if _, err := e.AttachFile(attach); err != nil {
			return err
		}
	}

	return e.SendWithStartTLS(fmt.Sprintf("%s:%s", facades.Config.GetString("mail.host"),
		facades.Config.GetString("mail.port")),
		LoginAuth(facades.Config.GetString("mail.username"),
			facades.Config.GetString("mail.password")), &tls.Config{ServerName: facades.Config.GetString("mail.host")})
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		}
	}
	return nil, nil
}
