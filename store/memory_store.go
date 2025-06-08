package store

import (
	"sort"
	"sync"
)

type URLStore struct {
	urlToShort map[string]string
	shortToURL map[string]string
	domainHits map[string]int
	mu         sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		urlToShort: make(map[string]string),
		shortToURL: make(map[string]string),
		domainHits: make(map[string]int),
	}
}

func (s *URLStore) GetShort(url string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	short, ok := s.urlToShort[url]
	return short, ok
}

func (s *URLStore) Save(url, short string, domain string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urlToShort[url] = short
	s.shortToURL[short] = url
	s.domainHits[domain]++
}

func (s *URLStore) GetOriginal(short string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.shortToURL[short]
	return url, ok
}

func (s *URLStore) GetTopDomains(n int) map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	type pair struct {
		domain string
		count  int
	}
	var pairs []pair
	for d, c := range s.domainHits {
		pairs = append(pairs, pair{d, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := make(map[string]int)
	for i := 0; i < n && i < len(pairs); i++ {
		result[pairs[i].domain] = pairs[i].count
	}
	return result
}
