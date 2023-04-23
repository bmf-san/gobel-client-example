package presenter

import (
	"html/template"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/model"
)

// ErrorData is a data for template.
type ErrorData struct {
	Code    int
	Message string
}

// Error responses a error template.
func (p *Presenter) Error(w http.ResponseWriter, code int) {
	e := &ErrorData{
		Code:    code,
		Message: handleErrorMessage(code),
	}

	fm := template.FuncMap{
		"year": p.year,
	}
	m := &model.Meta{
		Title:   "gobel-client-example.com - エラー",
		NoIndex: true,
	}
	tpl := template.Must(template.New("base").Funcs(fm).ParseFS(p.templates, "templates/layout/base.tpl", "templates/partial/meta.tpl", "templates/error/index.tpl"))

	w.WriteHeader(e.Code)

	if err := tpl.ExecuteTemplate(w, "base", map[string]interface{}{"Meta": m, "ErrorData": e}); err != nil {
		w.Write([]byte(err.Error()))
	}
}

// handleErrorMessage handles an error message by code.
func handleErrorMessage(code int) string {
	switch code {
	case http.StatusInternalServerError:
		return "すみません！エラーです！"
	}
	return ""
}
