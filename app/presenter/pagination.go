package presenter

import (
	"html/template"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// Pagination is a data for pagination template.
type Pagination struct {
	Pager       *model.Pagination
	QueryParams template.URL
}
