package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// CategoryIndex is a data for index template.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *model.Pagination
}

// ExecuteCategoryIndex responses a index template.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, c *CategoryIndex) error {
	tpl := template.Must(template.ParseFS(p.templates, "templates/layout/base.tpl", "templates/category/index.tpl", "templates/partial/pagination.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", c); err != nil {
		return err
	}
	return nil
}
