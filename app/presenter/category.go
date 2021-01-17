package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// CategoryIndex is a data for index view.
type CategoryIndex struct {
	Categories *model.Categories
	Pagination *model.Pagination
}

// ExecuteCategoryIndex responses a index view.
func (p *Presenter) ExecuteCategoryIndex(w http.ResponseWriter, c *CategoryIndex) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/category/index.tpl", "view/partial/pagination.tpl", "view/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", c); err != nil {
		return err
	}
	return nil
}
