package smtp

import (
	"crypto/tls"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type Service struct {
	cfg    Config
	mail   chan *gomail.Message
	dialer *gomail.Dialer
	conn   gomail.SendCloser
	opened bool
}

func NewService(c Config) *Service {
	return &Service{
		cfg:    c,
		mail:   make(chan *gomail.Message),
		dialer: nil,
		opened: false,
	}
}

func (s *Service) Open() error {
	go s.runMailer()
	return nil
}

func (s *Service) Close() error {
	close(s.mail)
	if s.opened {
		s.conn.Close()
	}

	return nil
}
func (s *Service) runMailer() {
	s.dialer = gomail.NewPlainDialer(s.cfg.Host, s.cfg.Port, s.cfg.Username, s.cfg.Password)
	if s.cfg.NoVerify {
		s.dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	var err error
	for {
		select {
		case m, ok := <-s.mail:
			if !ok {
				return
			}
			if !s.opened {
				if s.conn, err = s.dialer.Dial(); err != nil {
					log.Println("[Error] Dial Host error ", err)
					continue
				}
				s.opened = true
			}
			if err := gomail.Send(s.conn, m); err != nil {
				log.Println("[Error] Send Mail to ", m, " Error", err)
			}
		}
	}
}

func (s *Service) SendMail(subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.From)
	m.SetHeader("To", s.cfg.To...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	s.mail <- m
	return nil
}
