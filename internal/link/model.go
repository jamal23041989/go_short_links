package link

import (
	"github.com/jamal23041989/go_short_links/internal/stat"
	"gorm.io/gorm"
	"math/rand"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash(6)
	return link
}

func (link *Link) GenerateHash(n int) {
	link.Hash = RandStringRunes(n)
}

var letterRunes = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		b = append(b, letterRunes[rand.Intn(len(letterRunes))])
	}
	return string(b)
}
