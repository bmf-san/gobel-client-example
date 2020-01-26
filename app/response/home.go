package response

import (
	"html/template"
	"net/http"
)

// ExecuteHomeIndex responses a index view.
func (r *Response) ExecuteHomeIndex(w http.ResponseWriter) error {
	tpl := template.Must(template.ParseFiles("view/layout/base.tpl", "view/home/index.tpl"))
	if err := tpl.ExecuteTemplate(w, "base", ""); err != nil {
		return err
	}
	return nil
}
