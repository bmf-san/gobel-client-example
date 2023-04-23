package presenter

import (
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// ExecuteHomeIndex responses a index template.
func (pt *Presenter) ExecuteHomeIndex(w http.ResponseWriter, r *http.Request, p *PostIndex) error {
	fm := template.FuncMap{
		"year":      pt.year,
		"striptags": pt.StripTags,
		"summary":   pt.Summary,
		"isAd":      pt.IsAd,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(pt.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/home/index.tpl", "templates/partial/posts.tpl", "templates/partial/search.tpl"))
	u := os.Getenv("BASE_URL")
	m := &model.Meta{
		Title:         "gobel-client-example.com - ホーム",
		Canonical:     u,
		Description:   "gobel-client-example",
		OGTitle:       "gobel-client-example",
		OGDescription: "gobel-client-exampleはソフトウェアエンジニアであるbmf-sanが日々の技術ネタを投稿するサイトです。",
		OGURL:         u,
		OGType:        "website",
		OGImage:       "",
		OGSiteName:    "gobel-client-example",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Posts": p}); err != nil {
		return err
	}
	return nil
}
