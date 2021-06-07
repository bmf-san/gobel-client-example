package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// TagIndex is a data for index template.
type TagIndex struct {
	Tags       *model.Tags
	Pagination *model.Pagination
}

// ExecuteTagIndex responses a index template.
func (p *Presenter) ExecuteTagIndex(w http.ResponseWriter, t *TagIndex) error {
	tpl := template.Must(template.ParseFS(p.templates, "templates/layout/base.tpl", "templates/tag/index.tpl", "templates/partial/pagination.tpl", "templates/partial/pagination.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", t); err != nil {
		return err
	}
	return nil
}
