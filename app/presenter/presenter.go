package presenter

import (
	"embed"
	"html/template"
	"math/rand"
	"regexp"
	"time"
)

// Presenter represents the singular of presenter.
type Presenter struct {
	templates embed.FS
}

// NewPresenter creates a Presenter.
func NewPresenter(templates embed.FS) *Presenter {
	return &Presenter{
		templates: templates,
	}
}

// Unescape returns safe HTML.
func (p *Presenter) Unescape(text string) template.HTML {
	return template.HTML(text)
}

// year returns curernt year.
func (p *Presenter) year() int {
	return time.Now().Year()
}

const rgx string = `<.*?>`

// StripTags strips tags from string.
func (p *Presenter) StripTags(s string) string {
	r := regexp.MustCompile(rgx)
	return r.ReplaceAllString(s, "")
}

const defaultWords int = 200

// Summary returns summarized string.
func (p *Presenter) Summary(s string) string {
	r := []rune(s)
	var l int = defaultWords
	if len(r) < defaultWords {
		l = len(r)
	}
	return string(r[:l]) + "..."
}

// IsAd returns a flag indicating whether to output ads.
func (p *Presenter) IsAd() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(5)%3 == 0
}
