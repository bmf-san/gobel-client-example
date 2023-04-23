package presenter

import (
	"html/template"
	"net/http"
	"os"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// CategoryIndex is a data for index template.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *Pagination
}

// ExecuteCategoryIndex responses a index template.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, r *http.Request, c *CategoryIndex) error {
	fm := template.FuncMap{
		"year": p.year,
		"isAd": p.IsAd,
	}
	u := os.Getenv("BASE_URL") + "/categories"
	m := &model.Meta{
		Title:         "gobel-client-example.com - カテゴリ一覧",
		Canonical:     u,
		Description:   "カテゴリ一覧",
		OGTitle:       "カテゴリ一覧",
		OGDescription: "カテゴリ一覧",
		OGURL:         u,
		OGType:        "article",
		OGImage:       "",
		OGSiteName:    "gobel-client-example",
		OGLocale:      "ja_JP",
		TwitterCard:   "summary",
		TwitterSite:   "@bmf_san",
		NoIndex:       false,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/category/index.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "Categories": c}); err != nil {
		return err
	}
	return nil
}
