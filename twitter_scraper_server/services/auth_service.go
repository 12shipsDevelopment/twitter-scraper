package services

import (
	"twitter_scraper_server/config"

	twitterscraper "github.com/imperatrona/twitter-scraper"
)

type AuthService struct {
	cfg      *config.Config
	scrapers []*twitterscraper.Scraper
}

func NewAuthService(cfg *config.Config) *AuthService {
	scrapers := make([]*twitterscraper.Scraper, 0)
	scrapers = append(scrapers, twitterscraper.New())

	as := &AuthService{cfg: cfg, scrapers: scrapers}

	return as
}

func (as *AuthService) GetNextScraper() *twitterscraper.Scraper {
	return as.scrapers[0]
}
