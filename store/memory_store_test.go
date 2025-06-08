package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLStore_SaveAndGet(t *testing.T) {
	s := NewURLStore()
	url := "https://example.com"
	short := "abc123"
	s.Save(url, short, "example.com")

	gotShort, found := s.GetShort(url)
	assert.True(t, found)
	assert.Equal(t, short, gotShort)

	gotLong, found := s.GetOriginal(short)
	assert.True(t, found)
	assert.Equal(t, url, gotLong)
}

func TestURLStore_TopDomains(t *testing.T) {
	s := NewURLStore()
	s.Save("https://a.com/1", "a1", "a.com")
	s.Save("https://a.com/2", "a2", "a.com")
	s.Save("https://b.com/1", "b1", "b.com")
	s.Save("https://c.com/1", "c1", "c.com")
	s.Save("https://c.com/2", "c2", "c.com")
	s.Save("https://c.com/3", "c3", "c.com")

	top := s.GetTopDomains(2)

	assert.Equal(t, 3, top["c.com"])
	assert.Equal(t, 2, top["a.com"])
}
