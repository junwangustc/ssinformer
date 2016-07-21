package main

import (
	"log"

	"github.com/junwangustc/ssinformer/crawl"
	"github.com/junwangustc/ssinformer/smtp"
)

type Service interface {
	Open() error
	Close() error
}

type Server struct {
	smtpService  *smtp.Service
	crawlService *crawl.Service
	services     []Service
	exit         chan struct{}
}

func NewServer(c *Config) *Server {
	srv := &Server{
		exit: make(chan struct{}),
	}
	srv.smtpService = smtp.NewService(c.SMTP)
	srv.crawlService = crawl.NewService(c.Crawler)
	srv.services = append(srv.services, srv.smtpService)
	srv.services = append(srv.services, srv.crawlService)
	return srv
}
func (s *Server) Open() error {
	for _, srv := range s.services {
		if err := srv.Open(); err != nil {
			s.Close()
			return err
		}
	}
	return nil
}
func (s *Server) Close() error {
	for _, srv := range s.services {
		srv.Close()
	}
	close(s.exit)
	return nil
}
func (s *Server) StartCrawlAndSend() {

	log.Println("[Info] start Crawl And Send")
	for {
		select {
		case npsd := <-s.crawlService.GetNewPassword():
			s.smtpService.SendMail("shadowsocks password ", npsd)
			log.Println("[Info]Send Email over")
		case <-s.exit:
			log.Println("Exit the Crawl and Send")
			return

		}

	}

}
