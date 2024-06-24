package storage

import (
	"crypto/sha256"
	"fmt"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRanmdom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha256.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", err
	}
	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
