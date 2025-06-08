package service

import (
	"testing"

	"github.com/mohit83k/url-shortner/store"

	"github.com/stretchr/testify/assert"
)

func TestURLShortener_IdempotentShortening(t *testing.T) {
	s := NewURLShortener(store.NewURLStore())
	url := "https://www.youtube.com/watch?v=test"

	short1 := s.Shorten(url)
	short2 := s.Shorten(url)

	assert.Equal(t, short1, short2)

	long, ok := s.Resolve(short1)
	assert.True(t, ok)
	assert.Equal(t, url, long)
}

func TestURLShortener_TopDomains(t *testing.T) {
	s := NewURLShortener(store.NewURLStore())
	s.Shorten("https://youtube.com/1")
	s.Shorten("https://youtube.com/2")
	s.Shorten("https://wikipedia.org/1")

	top := s.TopDomains()

	assert.Equal(t, 2, top["youtube.com"])
	assert.Equal(t, 1, top["wikipedia.org"])
}
