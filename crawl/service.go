package crawl

import (
	"log"
	"time"

	"github.com/opesun/goquery"
)

type Service struct {
	crawlPassword string
	url           string
	interval      int
	msgChan       chan string
	exit          chan struct{}
}

func (s *Service) Open() error {
	go s.StartCrawl()
	return nil
}
func (s *Service) Close() error {
	close(s.exit)
	return nil

}
func NewService(c Config) *Service {
	return &Service{
		url:      c.URL,
		interval: c.Interval,
		msgChan:  make(chan string),
		exit:     make(chan struct{}),
	}
}

func (s *Service) StartCrawl() {

	ticker := time.NewTicker(time.Duration(s.interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			str := s.CrawlPSD()
			if str != s.crawlPassword {
				log.Println("[Info]Get New Password")
				s.crawlPassword = str
				s.msgChan <- s.crawlPassword
			} else {
				log.Println("[Info] Get the Same Password")
			}

		case <-s.exit:
			return

		}
	}

}
func (s *Service) GetNewPassword() chan string {
	return s.msgChan
}
func (s *Service) CrawlPSD() string {
	p, err := goquery.ParseUrl(s.url)
	if err != nil {
		log.Println(err)
		return s.crawlPassword
	}
	if len(p.Find("div.col-lg-4").Text()) > 731 {
		updatePassword := p.Find("div.col-lg-4").Text()[:730]
		return updatePassword
	}
	return s.crawlPassword
}
