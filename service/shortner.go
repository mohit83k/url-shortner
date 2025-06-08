package service

import (
	"net/url"
	"strings"

	"github.com/mohit83k/url-shortner/util"

	"github.com/mohit83k/url-shortner/store"
)

type URLShortener struct {
	Store *store.URLStore
}

func NewURLShortener(s *store.URLStore) *URLShortener {
	return &URLShortener{Store: s}
}

func (s *URLShortener) Shorten(longURL string) string {
	if short, found := s.Store.GetShort(longURL); found {
		return short
	}
	short := util.GenerateShortKey(longURL)
	domain := extractDomain(longURL)
	s.Store.Save(longURL, short, domain)
	return short
}

func (s *URLShortener) Resolve(short string) (string, bool) {
	return s.Store.GetOriginal(short)
}

func (s *URLShortener) TopDomains() map[string]int {
	return s.Store.GetTopDomains(3)
}

func extractDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	host := u.Host
	if strings.HasPrefix(host, "www.") {
		host = strings.TrimPrefix(host, "www.")
	}
	return host
}
